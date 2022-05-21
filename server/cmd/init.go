package main

import (
	"github.com/l1f/blockornot/internal/config"
	"github.com/urfave/cli/v2"
)

func init() {
	var command = cli.Command{
		Name:    "init",
		Usage:   "Generates a new config file at the specified path",
		Aliases: []string{"i", "install"},
		Action:  initConfig,
	}

	RegisterCommand(&command)
}

func initConfig(ctx *cli.Context) error {
	return config.GenerateNewConfigFile(ctx.String("config"))
}
