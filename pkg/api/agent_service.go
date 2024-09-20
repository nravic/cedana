package api

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/cedana/cedana/pkg/api/services/agent_task"
	"github.com/cedana/cedana/pkg/api/services/rpc"
	"github.com/cedana/cedana/pkg/api/services/task"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/mdlayher/vsock"
	"github.com/rs/zerolog"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthcheckgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
)

const VSOCK_PORT = 8080

type agentService struct {
	CRIU      *Criu
	logger    *zerolog.Logger
	serverCtx context.Context // context alive for the duration of the server
	wg        sync.WaitGroup  // for waiting for all background tasks to finish

	agent_task.UnimplementedTaskServiceServer
}

type AgentServer struct {
	grpcServer *grpc.Server
	service    *agentService
	listener   net.Listener
}

func NewAgentServer(ctx context.Context, opts *ServeOpts) (*AgentServer, error) {
	var err error

	machineID, err := utils.GetMachineID()
	if err != nil {
		return nil, err
	}

	server := &AgentServer{
		grpcServer: grpc.NewServer(
			grpc.StreamInterceptor(loggingStreamInterceptor()),
			grpc.UnaryInterceptor(loggingUnaryInterceptor(*opts, machineID)),
		),
	}

	healthcheck := health.NewServer()
	healthcheckgrpc.RegisterHealthServer(server.grpcServer, healthcheck)

	service := &agentService{
		// criu instantiated as empty, because all criu functions run criu swrk (starting the criu rpc server)
		// instead of leaving one running forever.
		CRIU:      &Criu{},
		serverCtx: ctx,
	}

	agent_task.RegisterTaskServiceServer(server.grpcServer, service)
	reflection.Register(server.grpcServer)

	var listener net.Listener

	listener, err = vsock.Listen(opts.Port, nil)

	if err != nil {
		return nil, err
	}
	server.listener = listener
	server.service = service

	healthcheck.SetServingStatus("task.TaskService", healthcheckgrpc.HealthCheckResponse_SERVING)
	return server, err
}

func (s *AgentServer) start(context.Context) error {
	return s.grpcServer.Serve(s.listener)
}

func (s *AgentServer) stop(context.Context) error {
	s.grpcServer.GracefulStop()
	return s.listener.Close()
}

func (s *agentService) prepareDumpOpts() *rpc.CriuOpts {
	opts := rpc.CriuOpts{
		LogLevel:     proto.Int32(CRIU_DUMP_LOG_LEVEL),
		LogFile:      proto.String(CRIU_DUMP_LOG_FILE),
		LeaveRunning: proto.Bool(viper.GetBool("client.leave_running")),
		GhostLimit:   proto.Uint32(GHOST_LIMIT),
	}
	return &opts
}

// prepareDump =/= preDump.
// prepareDump sets up the folders to dump into, and sets the criu options.
// preDump on the other hand does any process cleanup right before the checkpoint.
func (s *agentService) prepareDump(ctx context.Context, state *agent_task.ProcessState, args *agent_task.DumpArgs, opts *rpc.CriuOpts) (string, error) {
	stats, ok := ctx.Value("dumpStats").(*task.DumpStats)
	if !ok {
		return "", fmt.Errorf("could not get dump stats from context")
	}

	start := time.Now()

	var hasTCP bool
	var hasExtUnixSocket bool

	for _, Conn := range state.ProcessInfo.OpenConnections {
		if Conn.Type == syscall.SOCK_STREAM { // TCP
			hasTCP = true
		}

		if Conn.Type == syscall.AF_UNIX { // Interprocess
			hasExtUnixSocket = true
		}
	}

	opts.TcpEstablished = proto.Bool(hasTCP || args.TcpEstablished)
	opts.ExtUnixSk = proto.Bool(hasExtUnixSocket)
	opts.FileLocks = proto.Bool(true)

	// check tty state
	// if pts is in open fds, chances are it's a shell job
	var isShellJob bool
	for _, f := range state.ProcessInfo.OpenFds {
		if strings.Contains(f.Path, "pts") {
			isShellJob = true
			break
		}
	}
	opts.ShellJob = proto.Bool(isShellJob)
	opts.Stream = proto.Bool(args.Stream)

	// jobID + UTC time (nanoseconds)
	// strip out non posix-compliant characters from the jobID
	var dumpDirPath string
	if args.Stream {
		dumpDirPath = args.Dir
	} else {
		timeString := fmt.Sprintf("%d", time.Now().UTC().UnixNano())
		processDumpDir := strings.Join([]string{state.JID, timeString}, "_")
		dumpDirPath = filepath.Join(args.Dir, processDumpDir)
	}
	_, err := os.Stat(dumpDirPath)
	if err != nil {
		if err := os.MkdirAll(dumpDirPath, DUMP_FOLDER_PERMS); err != nil {
			return "", err
		}
	}

	err = os.Chown(args.Dir, int(state.UIDs[0]), int(state.GIDs[0]))
	if err != nil {
		return "", err
	}
	err = chownRecursive(dumpDirPath, state.UIDs[0], state.GIDs[0])
	if err != nil {
		return "", err
	}

	err = os.Chmod(args.Dir, DUMP_FOLDER_PERMS)
	if err != nil {
		return "", err
	}
	err = chmodRecursive(dumpDirPath, DUMP_FOLDER_PERMS)
	if err != nil {
		return "", err
	}

	// close common fds
	err = closeCommonFds(int32(os.Getpid()), state.PID)
	if err != nil {
		return "", err
	}

	elapsed := time.Since(start)
	stats.PrepareDuration = elapsed.Milliseconds()

	return dumpDirPath, nil
}

// TODO NR - customizable errors
func (s *agentService) generateState(ctx context.Context, pid int32) (*agent_task.ProcessState, error) {
	if pid == 0 {
		return nil, fmt.Errorf("invalid PID %d", pid)
	}

	state := &agent_task.ProcessState{}

	p, err := process.NewProcessWithContext(ctx, pid)
	if err != nil {
		return nil, fmt.Errorf("could not get gopsutil process: %v", err)
	}

	state.PID = pid

	// Search for JID, if found, use that state with existing fields
	// list, err := s.db.ListJobs(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not list jobs: %v", err)
	// }
	// for _, job := range list {
	// 	st := &task.ProcessState{}
	// 	err = json.Unmarshal(job.State, st)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	if st.PID == pid {
	// 		state = st
	// 		break
	// 	}
	// }

	var openFiles []*agent_task.OpenFilesStat
	var openConnections []*agent_task.ConnectionStat

	// get process uids, gids, and groups
	uids, err := p.UidsWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get uids: %v", err)
	}
	gids, err := p.GidsWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get gids: %v", err)
	}
	groups, err := p.GroupsWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get groups: %v", err)
	}
	state.UIDs = uids
	state.GIDs = gids
	state.Groups = groups

	of, err := p.OpenFiles()
	for _, f := range of {
		var mode string
		var stream agent_task.OpenFilesStat_StreamType
		file, err := os.Stat(f.Path)
		if err == nil {
			mode = file.Mode().Perm().String()
			switch f.Fd {
			case 0:
				stream = agent_task.OpenFilesStat_STDIN
			case 1:
				stream = agent_task.OpenFilesStat_STDOUT
			case 2:
				stream = agent_task.OpenFilesStat_STDERR
			default:
				stream = agent_task.OpenFilesStat_NONE
			}
		}

		openFiles = append(openFiles, &agent_task.OpenFilesStat{
			Fd:     f.Fd,
			Path:   f.Path,
			Mode:   mode,
			Stream: stream,
		})
	}

	if err != nil {
		// don't want to error out and break
		return nil, nil
	}
	// used for network barriers (TODO: NR)
	conns, err := p.Connections()
	if err != nil {
		return nil, nil
	}
	for _, conn := range conns {
		Laddr := &agent_task.Addr{
			IP:   conn.Laddr.IP,
			Port: conn.Laddr.Port,
		}
		Raddr := &agent_task.Addr{
			IP:   conn.Raddr.IP,
			Port: conn.Raddr.Port,
		}
		openConnections = append(openConnections, &agent_task.ConnectionStat{
			Fd:     conn.Fd,
			Family: conn.Family,
			Type:   conn.Type,
			Laddr:  Laddr,
			Raddr:  Raddr,
			Status: conn.Status,
			PID:    conn.Pid,
			UIDs:   conn.Uids,
		})
	}

	memoryUsed, _ := p.MemoryPercent()
	isRunning, _ := p.IsRunning()

	// if the process is actually running, we don't care that
	// we're potentially overriding a failed flag here.
	// In the case of a restored/resuscitated process this is a good thing

	// this is the status as returned by gopsutil.
	// ideally we want more than this, or some parsing to happen from this end
	status, _ := p.Status()

	// we need the cwd to ensure that it exists on the other side of the restore.
	// if it doesn't - we inheritFd it?
	cwd, err := p.Cwd()
	if err != nil {
		return nil, nil
	}

	// system information
	cpuinfo, _ := cpu.Info()
	vcpus, _ := cpu.Counts(true)
	state.CPUInfo = &agent_task.CPUInfo{
		Count:      int32(vcpus),
		CPU:        cpuinfo[0].CPU,
		VendorID:   cpuinfo[0].VendorID,
		Family:     cpuinfo[0].Family,
		PhysicalID: cpuinfo[0].PhysicalID,
	}

	mem, _ := mem.VirtualMemory()
	state.MemoryInfo = &agent_task.MemoryInfo{
		Total:     mem.Total,
		Available: mem.Available,
		Used:      mem.Used,
	}

	host, _ := host.Info()
	state.HostInfo = &agent_task.HostInfo{
		HostID:               host.HostID,
		Hostname:             host.Hostname,
		OS:                   host.OS,
		Platform:             host.Platform,
		KernelVersion:        host.KernelVersion,
		KernelArch:           host.KernelArch,
		VirtualizationSystem: host.VirtualizationSystem,
		VirtualizationRole:   host.VirtualizationRole,
	}

	// ignore sending network for now, little complicated
	state.ProcessInfo = &agent_task.ProcessInfo{
		OpenFds:         openFiles,
		WorkingDir:      cwd,
		MemoryPercent:   memoryUsed,
		IsRunning:       isRunning,
		OpenConnections: openConnections,
		Status:          strings.Join(status, ""),
	}

	return state, nil
}
