package trade

import (
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/locxiang/quantitative-trading/app/util"
	"github.com/lexkong/log"
	"strings"
	"time"
)

type Trade struct {
	models.BaseModel
	Symbol string `gorm:"not null;index:idx_symbol;type:varchar(20);" json:"symbol"`
	Type   string `json:"type"`
	AggTrade
}

type AggTrade struct {
	TradeID        int       `gorm:"column:trade_id" json:"trade_id"`
	Price          float64   `json:"price"`
	Quantity       float64   `json:"quantity"`
	FirstTradeID   int       `json:"first_trade_id"`
	LastTradeID    int       `json:"last_trade_id"`
	Timestamp      time.Time `gorm:"column:timestamp;not null;index:idx_symbol;" json:"timestamp"`
	BuyerMaker     bool      `json:"buyer_maker"`
	BestPriceMatch bool      `json:"best_price_match"`
}

//实现它的json序列化方法
func (t *AggTrade) MarshalJSON() int64 {
	var stamp = util.Time2Millisecond(t.Timestamp)
	return stamp
}

func (t *Trade) TableName() string {
	table := t.Symbol + "_trade"
	return strings.ToLower(table)
}

//数据迁移
func (t *Trade) Migrate() {

	exist := models.DB.HasTable(t)
	if exist {
		return
	}

	log.Infof("创建数据表：%s", t.TableName())

	//建表
	models.DB.AutoMigrate(t)

	t.Timestamp = time.Now()
	models.DB.Create(t)
}
