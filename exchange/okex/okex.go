package okex

import (
	"github.com/locxiang/GoEx/okcoin"

	"net/http"
	"os"
	"github.com/locxiang/GoEx"
	"github.com/go-ffmt/ffmt"
	"github.com/locxiang/quantitative-trading/app/models"
)

func Init() {
	os.Setenv("https_proxy", "socks5://127.0.0.1:8016")

	var okexFuture = okcoin.NewOKExSpot(http.DefaultClient, "", "")

	okexFuture.GetTickerWithWs(goex.EOS_USDT, func(ticker *goex.Ticker) {
		ffmt.Puts(ticker)

		okexFuture.GetExchangeName()

		t := new(models.Ticker)
		t.Exchange = "okex_spot"
		t.Symbol = ticker.Pair.String()
		t.Last = ticker.Last
		t.Buy = ticker.Buy
		t.Sell = ticker.Sell
		t.High = ticker.High
		t.Low = ticker.Low
		t.Vol = ticker.Vol
		t.Date = ticker.Date

		t.InsertTSDB()
	})


	select {

	}


	//okexFuture.ws.CloseWs()

}
