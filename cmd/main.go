package main

import (
	"code/gendiff"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args().Slice()
			if len(args) < 2 {
				return cli.Exit("Надо передать два параметра", 1)
			}

			diff, err := gendiff.GenDiff(args[0], args[1], cmd.String("format"))
			fmt.Println(diff)
			return err
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		panic(err)
	}
}
