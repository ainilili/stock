package cmd

import (
	"fmt"
	"github.com/ainilili/stock/proxy"
	"github.com/ainilili/stock/util/logger"
	"github.com/urfave/cli"
	"strconv"
	"strings"
)

func GetCommand(c *cli.Context) {
	if len(c.Args()) == 0 {
		return
	}
	query := c.Args()[0]

	stocks, err := proxy.List(query)
	if err != nil {
		logger.Error(err)
		return
	}
	if len(stocks) == 0 {
		logger.Error("没有查到任何内容")
		return
	}
	if len(stocks) == 1 {
		query = stocks[0].Code
	} else {
		logger.Infof("%-5s\t%-10s\t%s", "序号", "Name", "Code")
		for i, stock := range stocks {
			nameRune := []rune(stock.Name)
			if len(nameRune) > 10 {
				stock.Name = string(nameRune[:10])
			}
			logger.Infof("%-5s\t%-10s\t%s", strconv.Itoa(i+1)+".", stock.Name, strings.TrimSpace(stock.Code))
		}
		fmt.Printf("请输入序号选择股票: ")
		var index int
		_, err = fmt.Scan(&index)
		if err != nil {
			logger.Error(err)
			return
		}
		if index < 1 || index > len(stocks) {
			logger.Error("输入错误")
			return
		}
		query = stocks[index-1].Code
	}
	details, err := proxy.Get(query)
	if err != nil {
		logger.Errorf(err.Error())
		return
	}
	nameRune := []rune(details.Name)
	if len(nameRune) > 10 {
		details.Name = string(nameRune[:10])
	}
	logger.Infof("%-10s\t%-10s\t%-10s\t%-10s\t%-20s\t%-10s", "Name", "Code", "Prev Close", "Price", "Volume Transaction", "Change")
	logger.Infof("%-10s\t%-10s\t%-10s\t%-10s\t%-20s\t%-10s", details.Name, details.Code, details.PrevClose, details.Price, details.VolumeTransaction, details.Change+"%")

	logger.Infof("")
	logger.Infof(details.MinNewChart)
}
