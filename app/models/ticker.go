package models

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/locxiang/quantitative-trading/app/influxdb"
	"fmt"
	"github.com/locxiang/quantitative-trading/app/util"
)

type Ticker struct {
	BaseModel
	Exchange string  `json:"exchange"`
	Symbol   string  `gorm:"not null;index:idx_symbol;type:varchar(20);" json:"symbol"`
	Last     float64 `json:"last"` //最新价格
	Buy      float64 `json:"buy"`
	Sell     float64 `json:"sell"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Vol      float64 `json:"vol"`
	Date     uint64  `json:"date"` // 单位:秒(second)
}

func (t *Ticker) InsertTSDB() error {

	tags := map[string]string{
		"exchange": t.Exchange,
		"symbol":   t.Symbol,
	}

	fields := map[string]interface{}{
		"price": t.Last,
		"buy":   t.Buy,
		"sell":  t.Sell,
		"high":  t.High,
		"low":   t.Low,
		"vol":   t.Vol,
	}


	dateTime := util.Millisecond2Time(int64(t.Date))
	fmt.Printf("timestamp to datetime:%v\n", dateTime)

	pt, err := client.NewPoint("test_ticker", tags, fields, dateTime)
	if err != nil {
		return err
	}

	influxdb.PointsChan <- pt

	return nil

}
