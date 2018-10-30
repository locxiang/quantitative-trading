package pool

import (
	"time"
	"github.com/locxiang/quantitative-trading/app/models"
)

func GetTradePools() (pools []*TradePool) {
	for p := range disp.Pools {
		pools = append(pools, p)
	}
	return
}

//加载数据库中的所有任务
func LoadStartDB() {
	var pools []*models.Pool
	models.DB.Find(&pools)
	for _, pool := range pools {
		Register(Convert(pool))
	}
}

func Convert(p *models.Pool) *TradePool {
	tp := &TradePool{
		Id:       p.Id,
		Duration: p.Duration,
		Symbol:   p.Symbol,
	}

	return tp
}

//添加并且注册数据池
func Add(symbol string, d time.Duration) *TradePool {
	p := &models.Pool{
		Duration: d,
		Symbol:   symbol,
	}

	data := new(models.Pool)

	models.DB.FirstOrCreate(data, p)

	tp := Convert(data)
	Register(tp)
	return tp
}

//接收交易数据事件
func EventTrade(m *models.Trade) {
	disp.EventTrade <- m
}
