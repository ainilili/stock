package cmd

import (
	"github.com/urfave/cli"
)

var (
	Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "get stock information.",
			Action: GetCommand,
		},
		{
			Name:   "list",
			Usage:  "list of stocks.",
			Action: ListCommand,
		},
	}
)
