package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ainilili/stock/model"
	"github.com/ainilili/stock/util/http"
	"github.com/pterm/pterm"
	"strconv"
	"strings"
	"time"
)

type SinaProxy struct{}

func (p *SinaProxy) List(query string) ([]model.Stock, error) {
	resp, err := http.Get(fmt.Sprintf("https://suggest3.sinajs.cn/suggest/type=&key=%s&name=suggestdata_%d", query, time.Now().UnixNano()/1e6))
	if err != nil {
		return nil, err
	}
	body := parseSinaBody(string(resp))
	if body == "" {
		return nil, nil
	}
	untreatedInfos := strings.Split(body, ";")
	stocks := make([]model.Stock, 0)
	for _, untreatedInfo := range untreatedInfos {
		untreatedInfo = strings.TrimSpace(untreatedInfo)
		if untreatedInfo == "" {
			continue
		}
		parts := strings.Split(untreatedInfo, ",")
		stocks = append(stocks, model.Stock{
			Name: parts[0],
			Code: parts[3],
		})
	}
	return stocks, nil
}
func (p *SinaProxy) Get(query string) (*model.StockDetails, error) {
	resp, err := http.Get(fmt.Sprintf("https://hq.sinajs.cn/list=%s", query), http.HeaderOption{
		Name:  "Referer",
		Value: "https://finance.sina.com.cn/realstock/company/sh600000/nc.shtml",
	})
	if err != nil {
		return nil, err
	}
	body := parseSinaBody(string(resp))
	if body == "" || len([]rune(body)) < 20 {
		return nil, errors.New("response nil")
	}
	parts := strings.Split(body, ",")
	prevClose, _ := strconv.ParseFloat(parts[2], 64)
	price, _ := strconv.ParseFloat(parts[3], 64)
	volume, _ := strconv.ParseFloat(parts[8], 64)
	volumeTransaction, _ := strconv.ParseFloat(parts[9], 64)
	limitUp := fmt.Sprintf("%.2f", prevClose*1.1)
	limitDown := fmt.Sprintf("%.2f", prevClose*0.9)
	change := fmt.Sprintf("%.2f", (price-prevClose)/prevClose*100)
	changePrice := fmt.Sprintf("%.2f", price-prevClose)

	minChart, err := p.getNewChart(query)
	if err != nil {
		return nil, err
	}
	return &model.StockDetails{
		Name:              parts[0],
		Code:              query,
		PrevClose:         parts[2],
		Open:              parts[1],
		Price:             parts[3],
		Max:               parts[4],
		Min:               parts[5],
		LimitUp:           limitUp,
		LimitDown:         limitDown,
		Change:            change,
		ChangePrice:       changePrice,
		Volume:            fmt.Sprintf("%.2f", volume/(10000*100)),
		VolumeTransaction: fmt.Sprintf("%.2f", volumeTransaction/(10000*100)),
		MinNewChart:       minChart,
	}, nil
}

func (p *SinaProxy) getNewChart(query string) (string, error) {
	cookies, err := http.GetRespCookies("https://xueqiu.com")
	if err != nil {
		return "", err
	}
	xqat := ""
	for _, cookie := range cookies {
		if cookie.Name == "xqat" {
			xqat = cookie.Value
		}
	}
	body, err := http.Get(fmt.Sprintf("https://stock.xueqiu.com/v5/stock/chart/minute.json?symbol=%s&period=1d", strings.ToUpper(query)), http.HeaderOption{
		Name:  "cookie",
		Value: fmt.Sprintf("xqat=%s;", xqat),
	})
	if err != nil {
		return "", err
	}
	resp := new(model.StockChartResp)
	err = json.Unmarshal(body, resp)
	if err != nil {
		return "", err
	}
	positiveBars := pterm.Bars{}
	for _, item := range resp.Data.Items {
		style := pterm.NewStyle(pterm.FgGreen)
		if item.Current > resp.Data.LastClose {
			style = pterm.NewStyle(pterm.FgRed)
		}
		positiveBars = append(positiveBars, pterm.Bar{
			Label:      time.UnixMilli(item.Timestamp).Format("15:04") + "  " + fmt.Sprintf("%.2f", item.Current) + " " + fmt.Sprintf("%6.2f%%", (item.Current-resp.Data.LastClose)/resp.Data.LastClose*100),
			Value:      int((item.Current - resp.Data.LastClose) * 100),
			Style:      style,
			LabelStyle: style,
		})
	}
	_ = pterm.DefaultBarChart.
		WithHorizontal().
		WithBars(positiveBars).
		WithWidth(50).
		Render()
	return "", nil
}

func parseSinaBody(body string) string {
	body, _ = http.Decode(body, "gb18030")
	return strings.ReplaceAll(strings.Split(body, "=")[1], "\"", "")
}
