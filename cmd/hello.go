package cmd

import (
	"fmt"

	"github.com/WendellAdriel/go-simple-cli-boilerplate/core"
	"gopkg.in/urfave/cli.v1"
)

// Hello command
var Hello = cli.Command{
	Name:        "hello",
	Usage:       "Hello prints a message to the console",
	Description: "Hello prints a message to the console",
	Action:      hello,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "name, n",
			Usage:       "Name to print to the console",
			Value:       "world",
			Destination: &name,
		},
	},
}

func hello(c *cli.Context) error {
	author := core.AppConfig.GetString("app", "author")
	fmt.Printf("%s says: Hello, %s", author, name)

	return nil
}
