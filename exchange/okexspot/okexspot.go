package okexspot

import (
	"github.com/nntaoli-project/GoEx/okcoin"

	"net/http"
	"github.com/nntaoli-project/GoEx"
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/locxiang/quantitative-trading/exchange/util"
	"github.com/lexkong/log"
)

type Okex struct {
	AccessKey string
	SecretKey string
}

func (o *Okex) Init(tradeChan chan<- *models.Trade, pairs []goex.CurrencyPair) {

	var okExSpot = okcoin.NewOKExSpot(http.DefaultClient, o.AccessKey, o.SecretKey)

	for _, pair := range pairs {
		okExSpot.GetTradeWithWs(pair, func(trade *goex.Trade) {
			tradeChan <- util.ConvertTrade("okex.spot", trade)

		})
	}

	log.Info("okex.spot 建立成功")
}
