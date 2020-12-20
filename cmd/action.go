/**
2 * @Author: Nico
3 * @Date: 2020/12/20 20:10
4 */
package cmd

import (
	"github.com/ainilili/stock/logger"
	"github.com/ainilili/stock/proxy"
	"github.com/urfave/cli"
	"strings"
)

func GetCommand(c *cli.Context) {
	if len(c.Args()) == 0 {
		return
	}
	query := c.Args()[0]
	details, err := proxy.Get(query)
	if err != nil{
		return
	}
	nameRune := []rune(details.Name)
	if len(nameRune) > 10{
		details.Name = string(nameRune[:10])
	}
	logger.Infof("%-10s\t%-10s\t%-10s\t%-10s", "Name","Code","Price","Change")
	logger.Infof("%-10s\t%-10s\t%-10s\t%-10s", details.Name, details.Code, details.Price, details.Change + "%")
}

func ListCommand(c *cli.Context) {
	if len(c.Args()) == 0 {
		return
	}
	query := c.Args()[0]
	stocks, err := proxy.List(query)
	if err != nil{
		return
	}
	logger.Infof("%-10s\t%s", "Name","Code")
	for _, stock := range stocks{
		nameRune := []rune(stock.Name)
		if len(nameRune) > 10{
			stock.Name = string(nameRune[:10])
		}
		logger.Infof("%-10s\t%s", stock.Name, strings.TrimSpace(stock.Code))
	}
}
