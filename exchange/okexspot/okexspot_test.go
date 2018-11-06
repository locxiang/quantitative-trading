package okexspot_test

import (
	"testing"
	"github.com/nntaoli-project/GoEx/okcoin"
	"net/http"
	"github.com/nntaoli-project/GoEx"
	"os"
)

func TestInit(t *testing.T) {
	os.Setenv("https_proxy", "socks5://127.0.0.1:8016")

	var okExSpot = okcoin.NewOKExSpot(http.DefaultClient, "", "")

	okExSpot.GetTradeWithWs(goex.EOS_USDT, func(trade *goex.Trade) {
		//fmt.Printf("trade: %v \n",trade)

		//t := util.Millisecond2Time(trade.Date)
		//fmt.Printf("%s \n", t)

	})

	select {}
}
