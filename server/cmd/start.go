package main

import (
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"

	"github.com/l1f/blockornot/internal/application"
	"github.com/l1f/blockornot/internal/config"
	"github.com/l1f/blockornot/internal/logic"
	"github.com/l1f/blockornot/internal/server"
)

func init() {
	var command = cli.Command{
		Name:    "run",
		Usage:   "Starts the backend server of Block or Not",
		Aliases: []string{"start"},
		Action:  run,
	}

	RegisterCommand(&command)
}

func run(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger.Info().Msg("Initializing backend...")

	logger.Info().Msg("Loading configuration...")
	cfg, err := config.ReadFromFile(ctx.String("config"))
	if err != nil {
		logger.Fatal().Err(err)
	}

	appCtx := &application.Context{
		Config: *cfg,
		Logger: &logger,
	}
	appCtx.Logic = logic.New(appCtx)

	logger.Info().Msg("Stating webserver...")
	err = server.Start(appCtx)

	return err
}
