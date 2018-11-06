package exchange_test

import (
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
	"testing"
	"github.com/locxiang/quantitative-trading/exchange"
)

func init() {
	Initialization.ServerInit()
}

func TestStart(t *testing.T) {
	exchange.Start()

	select {}
}
