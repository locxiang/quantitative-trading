package huobi_test

import (
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
	"testing"
	"strings"
	"github.com/nntaoli-project/GoEx"
)

func init() {
	Initialization.ServerInit()
}

func TestM(t *testing.T) {

	ch := "market.eosusdt.detail"

	var currA, currB string
	if strings.HasSuffix(ch, "usdt.detail") {
		currB = "usdt"
	} else if strings.HasSuffix(ch, "husd.detail") {
		currB = "husd"
	} else if strings.HasSuffix(ch, "btc.detail") {
		currB = "btc"
	} else if strings.HasSuffix(ch, "eth.detail") {
		currB = "eth"
	} else if strings.HasSuffix(ch, "ht.detail") {
		currB = "ht"
	}

	currA = strings.TrimPrefix(ch, "market.")
	t.Log(ch,currA)
	currA = strings.TrimSuffix(currA, currB+".detail")

	a := goex.NewCurrency(currA, "")
	b := goex.NewCurrency(currB, "")
	pair := goex.NewCurrencyPair(a, b)

	t.Logf("%s ,%s ,%s", a, b, pair)

}
