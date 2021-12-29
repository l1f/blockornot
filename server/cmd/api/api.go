package main

import (
	"flag"
	"fmt"
	"github.com/l1f/blockornot/logger"
	"os"

	"github.com/l1f/blockornot/internal/application"
	"github.com/l1f/blockornot/internal/config"
	"github.com/l1f/blockornot/internal/logic"
	"github.com/l1f/blockornot/internal/server"
)

const version = "DEV"

type arguments struct {
	configPath string
}

var initCommand = flag.NewFlagSet("init", flag.ExitOnError)
var runCommand = flag.NewFlagSet("run", flag.ExitOnError)

var args arguments

func init() {
	flag.StringVar(&args.configPath, "c", "", fmt.Sprintf("The config path to the config file. "+
		"To generate a new config file, use %s -c <Path to the file to be generated> init", os.Args[0]))
	flag.Parse()

	if args.configPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(flag.Args()) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var err error
	switch flag.Args()[0] {
	case "init":
		err = initCommand.Parse(flag.Args()[1:])
	default:
		err = runCommand.Parse(flag.Args()[1:])
	}
	if err != nil {
		// Todo: Handle and print error
		initCommand.PrintDefaults()
		runCommand.PrintDefaults()
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	if initCommand.Parsed() {
		err := config.GenerateNewConfigFile(args.configPath)
		if err != nil {
			fmt.Println(err)
		}
	} else if runCommand.Parsed() {
		run()
	}
}

func run() {
	log := logger.Init()

	log.Info.Println("Initializing backend...")

	log.Info.Println("Loading configuration...")
	cfg, err := config.ReadFromFile(args.configPath)
	if err != nil {
		log.Error.Fatal(err)
	}

	ctx := &application.Context{
		Config: *cfg,
		Logger: &log,
	}
	ctx.Logic = logic.New(ctx)

	log.Info.Println("Stating webserver...")
	err = server.Start(ctx)
	if err != nil {
		log.Error.Fatal(err)
	}
}
