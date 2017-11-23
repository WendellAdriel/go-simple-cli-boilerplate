package main

import (
	"os"
	"runtime"

	"github.com/WendellAdriel/go-simple-cli-boilerplate/cmd"
	"github.com/WendellAdriel/go-simple-cli-boilerplate/core"
	"gopkg.in/urfave/cli.v1"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.App{
		Name:                 core.AppConfig.GetString("app", "name"),
		Usage:                core.AppConfig.GetString("app", "desc"),
		Version:              core.AppConfig.GetString("app", "version"),
		Author:               core.AppConfig.GetString("app", "author"),
		Email:                core.AppConfig.GetString("app", "email"),
		EnableBashCompletion: true,
		Commands: []cli.Command{
			cmd.Hello,
		},
	}

	app.Run(os.Args)
}
