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
			Name:    "transaction",
			Aliases: []string{"trans", "t"},
			Usage:   "commands to manipulate snapshots",
			Subcommands: []cli.Command{
				{
					Name:    "apply",
					Aliases: []string{"a"},
					Usage:   "Apply a transaction to the specified environment",
					Action:  notImplemented,
				},
				{
					Name:    "create",
					Aliases: []string{"c"},
					Usage:   "Create a transaction from the specified environment",
					Action:  notImplemented,
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "context",
			Value: "default",
			Usage: "Kubernetes context to operate on",
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
