package models

import "github.com/lexkong/log"

type Order struct {
	BaseModel
	BuyOrder
	SellOrder
	Profit     float64 `json:"profit"`      //利润
	Strategy   string  `json:"strategy"`    //策略名字
	StrategyId uint64  `json:"strategy_id"` //策略id
	PoolId     uint64  `json:"pool_id"`
}

type BuyOrder struct {
	BuyTradeId int64   `json:"buy_trade_id"` //买入订单id
	BuyPrice   float64 `json:"buy_price"`    //买入价格
	BuyTime    int64   `json:"buy_time"`     //买入时间
	Quantity   float64 `json:"quantity"`     //数量
}

type SellOrder struct {
	SellTradeId int64   `json:"sell_trade_id"` //卖出订单id
	SellPrice   float64 `json:"sell_price"`    //卖出价格
	SellTime    int64   `json:"sell_time"`     //卖出时间
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) Migrate() {
	exist := DB.HasTable(o)
	if exist {
		return
	}
	log.Infof("创建数据表：%s", o.TableName())

	//建表
	DB.AutoMigrate(o)
}
