/*
交易所模块

1. 根据配置加载要读取的币
2. 连接所有交易所
3. 获取所有Ticker数据，
4. 写入服务  service/ticker
 */
package exchange

import (
	"github.com/nntaoli-project/GoEx"
	"github.com/locxiang/quantitative-trading/app/pkg/setting"
	"github.com/lexkong/log"
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/locxiang/quantitative-trading/exchange/okexspot"
	"os"
	"github.com/locxiang/quantitative-trading/exchange/huobi"
)

//启动所有交易所
func Start() {
	os.Setenv("https_proxy", "socks5://127.0.0.1:8016")

	spots := setting.Env().Exchange.Spot
	ticker := make(chan *models.Ticker)
	trade := make(chan *models.Trade)

	go receiveMsg(ticker, trade)

	var pairs []goex.CurrencyPair
	for _, spot := range spots {
		pairs = append(pairs, goex.NewCurrencyPair2(spot))
	}

	//启动okex
	okex := &okexspot.Okex{
		AccessKey: "",
		SecretKey: "",
	}
	okex.Init(trade, pairs)

	//启动火币
	hb := &huobi.Okex{
	}
	hb.Init(trade, pairs)

}

//接收ticker 数据
func receiveMsg(tickerChan <-chan *models.Ticker, tradeChan chan *models.Trade) {
	for {
		select {
		case t, ok := <-tickerChan:
			if ok {
				t.InsertTSDB()
			} else {
				log.Error("tickerChan已经被关闭了", nil)
				return
			}
		case t, ok := <-tradeChan:
			if ok {
				t.InsertTSDB()
			} else {
				log.Error("tradeChan 已经被关闭了", nil)
				return
			}


		}
	}
}

type Exchange interface {
	Init(tradeChan chan<- *models.Trade, pairs []goex.CurrencyPair)
}
