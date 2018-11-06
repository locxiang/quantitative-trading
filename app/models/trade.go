package models

import (
	"github.com/lexkong/log"
	"strings"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/locxiang/quantitative-trading/app/util"
	"github.com/locxiang/quantitative-trading/app/influxdb"
)

type Trade struct {
	BaseModel
	Exchange      string  `json:"exchange"`
	Symbol        string  `gorm:"not null;index:idx_symbol;type:varchar(20);" json:"symbol"`
	TradeId       int64   `json:"trade_id"`
	Price         float64 `json:"price"`
	Quantity      float64 `json:"quantity"`
	IsMaker       bool    `json:"is_maker"`
	Ignore        bool    `json:"ignore"`
	BuyerOrderId  int64   `json:"buyer_order_id"`
	SellerOrderId int64   `json:"seller_order_id"`
	Direction     string  `json:"direction"`
	Timestamp     int64   `gorm:"column:timestamp;not null;index:idx_symbol;" json:"timestamp"`
}

const (
	OrderSell = "SELL"
	OrderBuy  = "BUY"
)


func (t *Trade) TableName() string {
	table := t.Symbol + "_trade"
	return strings.ToLower(table)
}

//数据迁移
func (t *Trade) Migrate() {

	exist := DB.HasTable(t)
	if exist {
		return
	}

	log.Infof("创建数据表：%s", t.TableName())

	//建表
	DB.AutoMigrate(t)
}


func (t *Trade) InsertTSDB() error {

	tags := map[string]string{
		"exchange":  t.Exchange,
		"symbol":    t.Symbol,
		"direction": t.Direction,
	}

	fields := map[string]interface{}{
		"trade_id": t.TradeId,
		"price":    t.Price,
		"quantity": t.Quantity,
	}

	var pts []*client.Point
	utcTime := util.Millisecond2Time(t.Timestamp)
	pt, err := client.NewPoint("test_trade", tags, fields, utcTime)
	if err != nil {
		fmt.Printf("NewPoint err %s", err)
		return err
	}
	pts = append(pts, pt)

	influxdb.PointsChan <- pt


	return nil

}
