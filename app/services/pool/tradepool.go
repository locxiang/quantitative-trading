package pool

import (
	"time"
	"sync"
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/locxiang/quantitative-trading/app/util"
	"fmt"
	"github.com/lexkong/log"
)

/**
交易数据池
此数据池实现三种能力
1. 交易明细
2. 交易量统计
3. 价格波动比率
 */
type TradePool struct {
	Id         uint64          `json:"id"`
	Duration   time.Duration   `json:"duration"` //在一定时间内
	Symbol     string          `json:"symbol"`   //类别
	Trades     []*models.Trade `json:"-"`
	SellCount  float64         `json:"sell_count"` //卖出量
	BuyCount   float64         `json:"buy_count"`  //买入量
	Ratio      float64         `json:"ratio"`      //涨幅百分比
	Strategies []Strategy      `json:"strategies"` //加载策略

	//TODO 调试方便
	FirstPriceP float64 `json:"first_price"`
	LastPriceP  float64 `json:"last_price"`
	sync.Mutex
}

type Strategy interface {
	Check(t *TradePool)
}

func (e *TradePool) AddStrategy(s Strategy) {
	defer e.Unlock()
	e.Lock()
	e.Strategies = append(e.Strategies, s)
}

func (e *TradePool) RemoveStrategy(s Strategy) {
	defer e.Unlock()
	e.Lock()

	for i, ss := range e.Strategies {
		if ss == s {
			log.Debugf("找到并删除")
			e.Strategies = append(e.Strategies[:i], e.Strategies[i+1:]...)
			return
		}
	}
}

func (e *TradePool) AcceptTrade(trade *models.Trade) {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Errorf(err.(error), "AcceptTrade recover")
		}
	}()
	//fmt.Printf("收到一条数据：%v , %s\n", trade, time.Now().Format("2006-01-02 15:04:05"))
	e.addAggTrade(trade)
	e.removeExpiredTrade()
	e.computeRatio()

	for _, s := range e.Strategies {
		s.Check(e)
	}

}

func (e *TradePool) GUID() string {

	return fmt.Sprintf("guid_%d", e.Id)
}

func (e *TradePool) GetSymbol() string {
	defer e.Unlock()
	e.Lock()
	return e.Symbol
}

func (e *TradePool) GetDuration() time.Duration {
	defer e.Unlock()
	e.Lock()
	return e.Duration
}

//最初价格
func (e *TradePool) FirstPrice() (float64) {
	defer e.Unlock()
	e.Lock()
	return e.Trades[0].Price
}

//最新价格
func (e *TradePool) LastPrice() (float64) {
	defer e.Unlock()
	e.Lock()
	l := len(e.Trades)
	if l == 0 {
		return -1
	}
	return e.Trades[l-1].Price
}

func (e *TradePool) addAggTrade(trade *models.Trade) {
	defer e.Unlock()
	e.Lock()

	if trade.Direction == models.OrderBuy {
		e.BuyCount += trade.Quantity
	} else {
		e.SellCount += trade.Quantity
	}
	e.Trades = append(e.Trades, trade)
	e.FirstPriceP = trade.Price
}

//过滤掉多余数据
func (e *TradePool) removeExpiredTrade() {
	defer e.Unlock()
	e.Lock()
	t1 := time.Now()

	if len(e.Trades) <= 1 {
		return
	}

	for {
		t2 := util.Millisecond2Time(e.Trades[0].Timestamp).Add(e.Duration)
		if t1.After(t2) {
			if e.Trades[0].Direction == models.OrderBuy {
				e.BuyCount -= e.Trades[0].Quantity
			} else {
				e.SellCount -= e.Trades[0].Quantity
			}
			e.Trades = e.Trades[1:]

			continue
		}
		break
	}
	e.LastPriceP = e.Trades[0].Price
}

// 获取此数据池的百分率
func (e *TradePool) computeRatio() {
	ratio := (e.LastPrice() - e.FirstPrice()) / e.FirstPrice()
	e.SetRatio(ratio)
}

func (e *TradePool) SetRatio(f float64) {
	defer e.Unlock()
	e.Lock()
	e.Ratio = f
}

func (e *TradePool) GetRatio() (f float64) {
	defer e.Unlock()
	e.Lock()
	return e.Ratio
}

func (e *TradePool) Close() error {
	panic("implement me")
}
