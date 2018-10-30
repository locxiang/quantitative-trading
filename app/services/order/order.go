package order

import (
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/locxiang/quantitative-trading/app/util"
	"time"
	"github.com/lexkong/log"
)

func Buy(poolid, strategyId uint64, strategyName string, quantity, price float64) *models.Order {
	order := new(models.Order)

	order.PoolId = poolid
	order.StrategyId = strategyId
	order.Strategy = strategyName
	order.BuyPrice = price
	order.Quantity = quantity
	order.BuyTime = util.Time2Millisecond(time.Now())

	models.DB.Create(order)

	return order
}

func Sell(order *models.Order, price float64) {
	update := new(models.Order)
	update.SellPrice = price
	update.SellTime = util.Time2Millisecond(time.Now())

	profit := (update.SellPrice * order.Quantity) - (order.BuyPrice * order.Quantity)
	profit -= (update.SellPrice * order.Quantity) * 0.0005  //购买手续费
	profit -= (order.BuyPrice * order.Quantity) * 0.0003    //出售手续费

	update.Profit = profit
	log.Debugf("%f*%f-%f*%f=%f", update.SellPrice, order.Quantity, order.BuyPrice, order.Quantity, order.Profit)

	models.DB.Model(order).Updates(update)
}

type Result struct {
	StrategyId uint64  `json:"strategy_id"` //策略id
	PoolId     uint64  `json:"pool_id"`
	ProfitSum  float64 `json:"profit_sum"`
	HitCount   uint64  `json:"hit_count"`
}

func ProfitList() (result []Result) {
	models.DB.Raw("SELECT sum(profit) as profit_sum,count(id) as hit_count ,strategy_id,pool_id FROM orders GROUP BY strategy_id,pool_id ORDER BY profit_sum DESC").Scan(&result)
	return
}
