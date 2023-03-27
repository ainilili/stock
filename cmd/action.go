package cmd

import (
	"github.com/ainilili/stock/proxy"
	"github.com/ainilili/stock/util/logger"
	"github.com/urfave/cli"
	"strings"
)

func GetCommand(c *cli.Context) {
	if len(c.Args()) == 0 {
		return
	}
	query := c.Args()[0]
	details, err := proxy.Get(query)
	if err != nil {
		logger.Errorf("get err: %v", err)
		return
	}
	nameRune := []rune(details.Name)
	if len(nameRune) > 10 {
		details.Name = string(nameRune[:10])
	}
	logger.Infof("%-10s\t%-10s\t%-10s\t%-20s\t%-10s", "Name", "Code", "Price", "Volume Transaction", "Change")
	logger.Infof("%-10s\t%-10s\t%-10s\t%-20s\t%-10s", details.Name, details.Code, details.Price, details.VolumeTransaction, details.Change+"%")

	logger.Infof("")
	logger.Infof(details.MinNewChart)
}

func ListCommand(c *cli.Context) {
	if len(c.Args()) == 0 {
		return
	}
	query := c.Args()[0]
	stocks, err := proxy.List(query)
	if err != nil {
		logger.Errorf("list err: %v", err)
		return
	}
	logger.Infof("%-10s\t%s", "Name", "Code")
	for _, stock := range stocks {
		nameRune := []rune(stock.Name)
		if len(nameRune) > 10 {
			stock.Name = string(nameRune[:10])
		}
		logger.Infof("%-10s\t%s", stock.Name, strings.TrimSpace(stock.Code))
	}
}
