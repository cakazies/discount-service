package main

import (
	"discount-service/api"
	"discount-service/infra"
	"os"

	"github.com/urfave/cli"
)

const (
	AppName = "Discount-Service"
	// AppTagLine Application tagline
	AppTagLine = "Discount Service"
)

func main() {

	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppTagLine
	app.Version = "0.0"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config,c",
			Value:  "config/config.toml",
			Usage:  "Main config",
			EnvVar: "APP_CONFIG_FILE",
		},
	}

	app.Commands = []cli.Command{
		API,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	_ = app.Run(os.Args)
}

// API run API server
var API = cli.Command{
	Name:     "api",
	Usage:    "Run API Server",
	HideHelp: true,
	Action: func(ctx *cli.Context) {
		api.NewServer(infra.New(ctx.GlobalString("config"))).Run()
	},
}
