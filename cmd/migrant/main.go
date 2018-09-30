package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func notImplemented(c *cli.Context) error {
	return fmt.Errorf("rot implemented")
}

func main() {
	app := cli.NewApp()
	app.Name = "migrant"
	app.Usage = "migrate among the clouds"
	app.Commands = []cli.Command{
		{
			Name:    "change",
			Aliases: []string{"c"},
			Usage:   "commands to manipulate change sets",
			Subcommands: []cli.Command{
				{
					Name:    "apply",
					Aliases: []string{"a"},
					Usage:   "Apply a transaction to the specified environment (implicitly sets the mark)",
					Action:  notImplemented,
				},
				{
					Name:    "create",
					Aliases: []string{"c"},
					Usage:   "Create a change set of all changes since the last mark",
					Action:  notImplemented,
				},
				{
					Name:    "mark",
					Aliases: []string{"m"},
					Usage:   "Mark the start of a new change.",
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "context",
			Value: "default",
			Usage: "Kubernetes context to operate on",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
