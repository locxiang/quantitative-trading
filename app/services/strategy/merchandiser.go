package strategy

import (
	"github.com/locxiang/quantitative-trading/app/services/pool"
	"github.com/lexkong/log"
	orderService "github.com/locxiang/quantitative-trading/app/services/order"
	"time"
	"github.com/locxiang/quantitative-trading/app/models"
	"fmt"
	"github.com/locxiang/quantitative-trading/app/util"
)

//跟单策略
type Merchandiser struct {
	Id          uint64  //策略id
	Ratio       float64 //涨幅比例
	BuyQuantity float64 //购买数量
	GoodRatio   float64 //止盈比例
	BadRatio    float64 //止损比例
	buyPrice    float64 //购买价格
	pool        *pool.TradePool
}

//策略写入数据库
func (s *Merchandiser) Insert() (strategyId uint64) {
	strategy := new(models.Strategy)

	sm := &models.StrategyMerchandiser{
		Ratio:       s.Ratio,
		BuyQuantity: s.BuyQuantity,
		GoodRatio:   s.GoodRatio,
		BadRatio:    s.BadRatio,
	}
	strategy.Type = sm.Type()
	strategy.StrategyMerchandiser = sm

	//如果不存在就创建
	models.DB.FirstOrCreate(strategy, strategy)

	strategyId = strategy.Id
	return
}

func (s *Merchandiser) Name() string {
	return fmt.Sprintf("%f,%f,%f,%f", s.Ratio, s.BuyQuantity, s.GoodRatio, s.BadRatio)
}

//止盈价格
func (s *Merchandiser) GoodPrice() float64 {
	n := s.buyPrice * (1 + s.GoodRatio)
	return util.Round(n, 4)
}

//止损价格
func (s *Merchandiser) BadPrice() float64 {
	n := s.buyPrice * (1 + s.BadRatio)
	return util.Round(n, 4)
}

//检查策略如果命中就执行任务
func (s *Merchandiser) Check(tradePool *pool.TradePool) {
	s.pool = tradePool
	if s.Ratio > 0 && tradePool.GetRatio() >= s.Ratio {
		log.Debugf("命中策略:Merchandiser, %f >= %f", tradePool.GetRatio(), s.Ratio)
	} else if s.Ratio < 0 && tradePool.GetRatio() <= s.Ratio {
		log.Debugf("命中策略:Merchandiser, %f <= %f", tradePool.GetRatio(), s.Ratio)
	} else {
		return
	}

	//如果命中就购买
	s.buyPrice = tradePool.FirstPrice()
	order := orderService.Buy(tradePool.Id, s.Id, s.Name(), s.BuyQuantity, s.buyPrice)

	//删除此策略
	tradePool.RemoveStrategy(s)

	go func() {
		s.monitoring(order)
		//添加策略
		tradePool.AddStrategy(s)
	}()

}

//TODO 临时的，以后改成真实交易就不用了
//监听数据变化，准备卖出
func (s *Merchandiser) monitoring(order *models.Order) {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Errorf(err.(error), "AcceptTrade recover: %+v", s.pool.Trades)
		}
	}()

	for {
		if s.pool.FirstPrice() >= s.GoodPrice() {
			log.Infof("止盈卖出")
			orderService.Sell(order, s.GoodPrice())
			return

		} else if s.pool.FirstPrice() <= s.BadPrice() {
			log.Warnf("止损卖出")
			orderService.Sell(order, s.BadPrice())
			return
		}

		time.Sleep(300 * time.Millisecond)
	}

}
