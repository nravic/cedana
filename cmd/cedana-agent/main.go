package main

import (
	"context"
	"os"
	"time"

	"github.com/cedana/cedana/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var cudaVersions = map[string]string{
	"11.8": "cuda11_8",
	"12.1": "cuda12_1",
	"12.2": "cuda12_2",
	"12.4": "cuda12_4",
}

func main() {
	ctx := context.Background()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	if os.Getuid() != 0 {
		log.Error().Msg("daemon must be run as root")
	}

	// _, err := utils.InitOtel(cmd.Context(), cmd.Parent().Version)
	// if err != nil {
	// 	logger.Warn().Err(err).Msg("Failed to initialize otel")
	// }

	// logger.Info().Msg("otel initialized")

	// if viper.GetBool("profiling_enabled") {
	// 	go startProfiler()
	// }

	srvOpts := &api.ServeOpts{}

	server, err := api.NewAgentServer(ctx, srvOpts)
	if err != nil {
		log.Err(err)
	}

	log.Logger.Debug().Interface("Server:", server)

	err = api.StartServer(ctx, srvOpts, server)
	if err != nil {
		log.Err(err)
	}
}
