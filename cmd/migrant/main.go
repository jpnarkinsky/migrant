package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jpnarkinsky/migrant"
	"github.com/urfave/cli"
)

func notImplemented(c *cli.Context) error {
	return fmt.Errorf("rot implemented")
}

func main() {
	app := cli.NewApp()
	app.Name = "migrant"
	app.Version = migrant.Version
	app.Usage = "migrate among the clouds"
	app.Commands = []cli.Command{
		{
			Name:    "change",
			Aliases: []string{"c"},
			Usage:   "commands to manipulate change sets",
			Subcommands: []cli.Command{
				{
					Name:   "apply",
					Usage:  "Apply a transaction to the specified environment (sets the mark after the change is applied)",
					Action: notImplemented,
				},
				{
					Name:   "mark",
					Usage:  "Mark the start of a new change.",
					Action: notImplemented,
				},
				{
					Name:   "save",
					Usage:  "Save the current change",
					Action: notImplemented,
				},
				{
					Name:   "status",
					Usage:  "See the current status of the changeset",
					Action: notImplemented,
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
