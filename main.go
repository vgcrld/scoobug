package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "scoobug",
		Usage: "A simple CLI app example",
		Commands: []*cli.Command{
			{
				Name:    "greet",
				Aliases: []string{"gr"},
				Usage:   "Prints a greeting message",
				Action: func(c *cli.Context) error {
					fmt.Println("Hello, welcome to scoobug CLI!")
					return nil
				},
			},
			{
				Name:    "leave",
				Aliases: []string{"le"},
				Usage:   "I'm out of here!",
				Action: func(c *cli.Context) error {
					fmt.Println("It's time to rolllllll.")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
