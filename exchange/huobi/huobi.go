package huobi

import (
	"github.com/locxiang/GoEx/huobi"
	"net/http"
	"github.com/nntaoli-project/GoEx"
	"github.com/locxiang/quantitative-trading/app/models"
	"os"
	"net/url"
	"time"
	"net"
)

func Init() {

	os.Setenv("https_proxy", "socks5://127.0.0.1:8016")

	var httpProxyClient = &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return &url.URL{
					Scheme: "socks5",
					Host:   "127.0.0.1:8016"}, nil
			},
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
		},
		Timeout: 10 * time.Second,
	}

	var huoBiPro = huobi.NewHuoBiProSpot(httpProxyClient, "", "")

	huoBiPro.GetTickerWithWs(goex.EOS_USDT, func(ticker *goex.Ticker) {

		t := new(models.Ticker)
		t.Exchange = huoBiPro.GetExchangeName()
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

}
