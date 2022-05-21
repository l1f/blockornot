package main

import (
	"github.com/rs/zerolog"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "DEV"

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
var commands cli.Commands

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func RegisterCommand(command *cli.Command) {
	commands = append(commands, command)
}

func main() {
	app := &cli.App{
		Name:     "Block or Not",
		HelpName: "- block twitter users..",
		Version:  version,
		Commands: commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "The config path to the config file.",
				EnvVars:  []string{"BON_CONFIG_PATH"},
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Runs the application in debug mode.",
				EnvVars: []string{"BON_DEBUG"},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal().Err(err)
	}
}
