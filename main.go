package main

import (
	"fmt"
	"github.com/ainilili/stock/cmd"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

func main() {
	app := cli.NewApp()
	app.Usage = "a stock tools for you."
	app.Version = fmt.Sprintf("stock %s %s/%s", "v1.0.0", runtime.GOOS, runtime.GOARCH)
	app.Commands = cmd.Commands
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
	}
}
