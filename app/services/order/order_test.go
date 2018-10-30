package order

import (
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
	"testing"
	"github.com/go-ffmt/ffmt"
)

func init() {
	Initialization.ServerInit()
}

func TestBuy(t *testing.T) {
	order := Buy(3, 3, "test", 100, 500)
	Sell(order, 600)
}

func TestProfitList(t *testing.T) {
	s := ProfitList()
	ffmt.Puts(s)
}