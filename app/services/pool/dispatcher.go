package pool

import (
	"github.com/locxiang/quantitative-trading/app/models"
	"sync"
)

var disp *Dispatcher

func Init() {
	disp = newDispatcher()
	go disp.Run()

	//加载数据库中的策略数据池
	LoadStartDB()
}

//注册数据池
func Register(t *TradePool)  {
	disp.Register <- t
}

func Unregister(t *TradePool) {
	disp.Unregister <- t
}

type Dispatcher struct {
	// Registered TradePool.
	Pools      map[*TradePool]bool
	EventTrade chan *models.Trade
	Register   chan *TradePool
	Unregister chan *TradePool
	sync.Mutex
}


func newDispatcher() *Dispatcher {
	return &Dispatcher{
		EventTrade: make(chan *models.Trade),
		Register:   make(chan *TradePool),
		Unregister: make(chan *TradePool),
		Pools:      make(map[*TradePool]bool),
	}
}

func (d *Dispatcher) Run() {
	for {
		select {
		case pool := <-d.Register:
			d.Pools[pool] = true
		case pool := <-d.Unregister:
			if _, ok := d.Pools[pool]; ok {
				delete(d.Pools, pool)
				pool.Close()
			}
		case m := <-d.EventTrade:
			for pool := range d.Pools {
				if pool.GetSymbol() == m.Symbol {
					pool.AcceptTrade(m)
				}
			}
		}
	}
}
