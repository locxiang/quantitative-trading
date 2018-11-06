package util

import (
	"github.com/nntaoli-project/GoEx"
	"github.com/locxiang/quantitative-trading/app/models"
)

func ConvertTicker(exchangeName string, ticker *goex.Ticker) *models.Ticker {
	t := new(models.Ticker)
	t.Exchange = exchangeName
	t.Symbol = ticker.Pair.String()
	t.Last = ticker.Last
	t.Buy = ticker.Buy
	t.Sell = ticker.Sell
	t.High = ticker.High
	t.Low = ticker.Low
	t.Vol = ticker.Vol
	t.Date = ticker.Date

	t.InsertTSDB()
	return t
}

func ConvertTrade(exchangeName string, trade *goex.Trade) *models.Trade {
	t := new(models.Trade)
	t.Exchange = exchangeName
	t.Symbol = trade.Pair.String()
	t.TradeId = trade.Tid
	t.Price = trade.Price
	t.Quantity = trade.Amount
	t.Direction = trade.Type.String()
	t.Timestamp = trade.Date


	return t
}
