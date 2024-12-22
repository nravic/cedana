package cmd

import (
	"fmt"
	"os"
	"strings"

	"buf.build/gen/go/cedana/cedana/protocolbuffers/go/daemon"
	"github.com/cedana/cedana/internal/server"
	"github.com/cedana/cedana/pkg/config"
	"github.com/cedana/cedana/pkg/flags"
	"github.com/cedana/cedana/pkg/style"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	daemonCmd.AddCommand(startDaemonCmd)
	daemonCmd.AddCommand(checkDaemonCmd)

	// Add flags
	startDaemonCmd.PersistentFlags().
		BoolP(flags.MetricsASRFlag.Full, flags.MetricsASRFlag.Short, false, "enable metrics for ASR")
	startDaemonCmd.PersistentFlags().
		StringP(flags.DBFlag.Full, flags.DBFlag.Short, "", "path to local database")
	checkDaemonCmd.PersistentFlags().
		BoolP(flags.FullFlag.Full, flags.FullFlag.Short, false, "perform a full check (including plugins)")

	// Bind to config
	viper.BindPFlag("metrics.asr", startDaemonCmd.PersistentFlags().Lookup(flags.MetricsASRFlag.Full))
	viper.BindPFlag("db.path", startDaemonCmd.PersistentFlags().Lookup(flags.DBFlag.Full))
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage the daemon",
}

var startDaemonCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		if utils.IsRootUser() == false {
			return fmt.Errorf("daemon must be run as root")
		}

		var err error

		log.Info().Str("version", rootCmd.Version).Msg("starting daemon")

		server, err := server.NewServer(cmd.Context(), &server.ServeOpts{
			UseVSOCK: config.Global.UseVSOCK,
			Port:     config.Global.Port,
			Host:     config.Global.Host,
			Metrics:  config.Global.Metrics,
			Version:  cmd.Version,
		})
		if err != nil {
			log.Error().Err(err).Msgf("stopping daemon")
			return fmt.Errorf("failed to create server: %w", err)
		}

		err = server.Launch(cmd.Context())
		if err != nil {
			log.Error().Err(err).Msgf("stopping daemon")
			return err
		}

		return nil
	},
}

var checkDaemonCmd = &cobra.Command{
	Use:   "check",
	Short: "Health check the daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		useVSOCK := config.Global.UseVSOCK
		var client *Client
		var err error

		if useVSOCK {
			client, err = NewVSOCKClient(config.Global.ContextID, config.Global.Port)
		} else {
			client, err = NewClient(config.Global.Host, config.Global.Port)
		}
		if err != nil {
			return fmt.Errorf("Error creating client: %v", err)
		}

		full, _ := cmd.Flags().GetBool(flags.FullFlag.Full)

		resp, err := client.HealthCheck(cmd.Context(), &daemon.HealthCheckReq{Full: full})
		if err != nil {
			return err
		}

		results := resp.GetResults()

		errorCount := 0
		warningCount := 0

		tableWriter := table.NewWriter()
		tableWriter.SetStyle(style.TableStyle)
		tableWriter.SetOutputMirror(os.Stdout)

		for _, result := range results {
			tableWriter.AppendSeparator()
			name := strings.ToUpper(result.Name)
			tableWriter.AppendRow(table.Row{text.Bold.Sprint(name), "", ""})
			for _, component := range result.Components {
				statusErr := style.NegativeColor.Sprint(style.CrossMark)
				statusWarn := style.WarningColor.Sprint(style.BulletMark)
				statusOk := style.PositiveColor.Sprint(style.TickMark)
				data := component.Data
				var status string
				if len(component.Errors) > 0 {
					status = statusErr
					data = style.NegativeColor.Sprint(data)
				} else if len(component.Warnings) > 0 {
					status = statusWarn
					data = style.WarningColor.Sprint(data)
				} else {
					status = statusOk
				}

				maxLinelen := 60
				rows := []table.Row{{component.Name, data, status}}
				for _, err := range component.Errors {
					errorCount++
					err = style.BreakLine(err, maxLinelen)
					err = style.DisabledColor.Sprint(err)
					if len(rows) == 1 && len(rows[0]) == 3 {
						rows[0] = append(rows[0], err)
						continue
					}
					rows = append(rows, table.Row{"", "", statusErr, err})
				}
				for _, warn := range component.Warnings {
					warningCount++
					warn = style.BreakLine(warn, maxLinelen)
					warn = style.DisabledColor.Sprint(warn)
					if len(rows) == 1 && len(rows[0]) == 3 {
						rows[0] = append(rows[0], warn)
						continue
					}
					rows = append(rows, table.Row{"", "", statusWarn, warn})
				}

				tableWriter.AppendRows(rows)
			}
		}

		tableWriter.Render()
		fmt.Println()

		if errorCount > 0 {
			if warningCount > 0 {
				return fmt.Errorf("Failed with %d error(s) and %d warning(s).", errorCount, warningCount)
			}
			return fmt.Errorf("Failed with %d error(s).", errorCount)
		} else if warningCount > 0 {
			fmt.Printf("Looks good, with %d warning(s).\n", warningCount)
		} else {
			fmt.Println("All good.")
		}

		return nil
	},
}
