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
	MinNewChart       string //分时走势图
}

type StockChartItem struct {
	Current       float64                     `json:"current"`
	Volume        float64                     `json:"volume"`
	AvgPrice      float64                     `json:"avg_price"`
	Chg           float64                     `json:"chg"`
	Percent       float64                     `json:"percent"`
	Timestamp     int64                       `json:"timestamp"`
	Amount        float64                     `json:"amount"`
	High          float64                     `json:"high"`
	Low           float64                     `json:"low"`
	Macd          interface{}                 `json:"macd"`
	Kdj           interface{}                 `json:"kdj"`
	Ratio         interface{}                 `json:"ratio"`
	Capital       StockChartItemCapital       `json:"capital"`
	VolumeCompare StockChartItemVolumeCompare `json:"volume_compare"`
}

type StockChartItemCapital struct {
	Small  float64 `json:"small"`
	Medium float64 `json:"medium"`
	Large  float64 `json:"large"`
	Xlarge float64 `json:"xlarge"`
}

type StockChartItemVolumeCompare struct {
	VolumeSum     int `json:"volume_sum"`
	VolumeSumLast int `json:"volume_sum_last"`
}

type StockChart struct {
	LastClose float64          `json:"last_close"`
	After     []interface{}    `json:"after"`
	Items     []StockChartItem `json:"items"`
	ItemsSize int              `json:"items_size"`
}

type StockChartResp struct {
	Data             StockChart `json:"data"`
	ErrorCode        int        `json:"error_code"`
	ErrorDescription string     `json:"error_description"`
}
