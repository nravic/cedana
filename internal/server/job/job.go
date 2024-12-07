package job

// Job is a thread-safe wrapper for proto Job that is used to represent a managed process, or container.
// Simply embeds the proto Job struct, allowing us to add thread-safe methods to it.
// Allows multiple concurrent readers, but only one concurrent writer.

import (
	"os"
	"sync"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	"github.com/cedana/cedana/pkg/criu"
	"github.com/shirou/gopsutil/v4/process"
)

type Job struct {
	JID   string
	proto daemon.Job

	// Notify callbacks that can be saved for later use.
	// Will be called each time the job is C/R'd.
	CRIUCallback *criu.NotifyCallback

	sync.RWMutex
}

func newJob(
	jid string,
	jobType string,
) *Job {
	return &Job{
		JID: jid,
		proto: daemon.Job{
			JID:  jid,
			Type: jobType,
		},
		CRIUCallback: &criu.NotifyCallback{},
	}
}

func fromProto(j *daemon.Job) *Job {
	return &Job{
		JID: j.GetJID(),
		proto: daemon.Job{
			JID:            j.GetJID(),
			Type:           j.GetType(),
			Process:        j.GetProcess(),
			Details:        j.GetDetails(),
			Log:            j.GetLog(),
			CheckpointPath: j.GetCheckpointPath(),
			GPUEnabled:     j.GetGPUEnabled(),
		},
		CRIUCallback: &criu.NotifyCallback{},
	}
}

func (j *Job) GetPID() uint32 {
	j.RLock()
	defer j.RUnlock()
	return j.proto.GetProcess().GetPID()
}

func (j *Job) SetPID(pid uint32) {
	j.Lock()
	defer j.Unlock()
	if j.proto.GetProcess() == nil {
		j.proto.Process = &daemon.ProcessState{}
	}
	if j.proto.GetProcess().GetInfo() == nil {
		j.proto.Process.Info = &daemon.ProcessInfo{}
	}
	j.proto.Process.PID = pid
	j.proto.Process.Info.PID = pid
}

func (j *Job) GetProto() *daemon.Job {
	j.RLock()
	defer j.RUnlock()

	// Get all latest info
	j.proto.Process = j.GetProcess()
	j.proto.Log = j.GetLog()

	return &j.proto
}

func (j *Job) GetType() string {
	j.RLock()
	defer j.RUnlock()
	return j.proto.Type
}

func (j *Job) SetType(jobType string) {
	j.Lock()
	defer j.Unlock()
	j.proto.Type = jobType
}

func (j *Job) GetProcess() *daemon.ProcessState {
	j.RLock()
	defer j.RUnlock()

	pid := j.GetPID()

	if j.proto.GetProcess() == nil {
		j.proto.Process = &daemon.ProcessState{}
	}
	if j.proto.GetProcess().GetInfo() == nil {
		j.proto.Process.Info = &daemon.ProcessInfo{}
	}

	j.proto.Process.Info.Status = "halted"
	j.proto.Process.Info.IsRunning = false

	p, err := process.NewProcess(int32(pid))
	if err == nil {
		status, err := p.Status()
		if err == nil {
			j.proto.Process.Info.Status = status[0]
			j.proto.Process.Info.IsRunning = true
		}
	}

	return j.proto.Process
}

func (j *Job) SetProcess(process *daemon.ProcessState) {
	j.Lock()
	defer j.Unlock()
	j.proto.Process = process
}

func (j *Job) GetDetails() *daemon.Details {
	j.RLock()
	defer j.RUnlock()
	return j.proto.Details
}

func (j *Job) SetDetails(details *daemon.Details) {
	j.Lock()
	defer j.Unlock()
	j.proto.Details = details
}

func (j *Job) GetLog() string {
	j.RLock()
	defer j.RUnlock()

	// Check if log file exists
	log := j.proto.Log
	if _, e := os.Stat(log); os.IsNotExist(e) {
		return ""
	}

	return log
}

func (j *Job) SetLog(log string) {
	j.Lock()
	defer j.Unlock()
	j.proto.Log = log
}

func (j *Job) SetCheckpointPath(path string) {
	j.Lock()
	defer j.Unlock()
	j.proto.CheckpointPath = path
}

func (j *Job) GetCheckpointPath() string {
	j.RLock()
	defer j.RUnlock()
	return j.proto.CheckpointPath
}

func (j *Job) IsRunning() bool {
	return j.GetProcess().GetInfo().GetIsRunning()
}

func (j *Job) GPUEnabled() bool {
	j.RLock()
	defer j.RUnlock()
	return j.proto.GPUEnabled
}

func (j *Job) SetGPUEnabled(enabled bool) {
	j.Lock()
	defer j.Unlock()
	j.proto.GPUEnabled = enabled
}
