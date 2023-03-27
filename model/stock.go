package model

type Stock struct {
	Name string
	Code string
}

type StockDetails struct {
	Name              string //名称
	Code              string //code
	PrevClose         string //昨收
	Open              string //开盘
	Price             string //现价
	Max               string //最高
	Min               string //最低
	LimitUp           string //涨停
	LimitDown         string //跌停
	Change            string //涨幅
	ChangePrice       string //涨跌
	Volume            string //成交量
	VolumeTransaction string //成交额
	MarketCap         string //市值
	TurnoverRate      string //换手率
	PERatio           string //市盈率
	PBRatio           string //市净率
	EPS               string //每股收益
}
