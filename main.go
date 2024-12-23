package main

import (
	"fmt"

	// You can use _ to import a package without using it
	// Then you can init stuff in the package with func init()
	_ "github.com/vgcrld/scoobug/cfg"
)

func main() {

	fmt.Println("test of main")

	// app := &cli.App{
	// 	Name:  "scoobug",
	// 	Usage: "A simple CLI app example",
	// 	Commands: []*cli.Command{
	// 		{
	// 			Name:    "greet",
	// 			Aliases: []string{"gr"},
	// 			Usage:   "Prints a greeting message",
	// 			Action: func(c *cli.Context) error {
	// 				person := cfg.Person{}
	// 				person.SetName(c.String("name"))
	// 				fmt.Printf("Hello, %s!\n", person.GetName())
	// 				return nil
	// 			},
	// 			Flags: []cli.Flag{
	// 				&cli.StringFlag{
	// 					Name:    "name",
	// 					Aliases: []string{"n"},
	// 					Value:   "stranger",
	// 					Usage:   "your name",
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name:    "leave",
	// 			Aliases: []string{"le"},
	// 			Usage:   "I'm out of here!",
	// 			Action: func(c *cli.Context) error {
	// 				fmt.Println("It's time to rolllllll.")
	// 				return nil
	// 			},
	// 		},
	// 	},
	// }

	// err := app.Run(os.Args)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
