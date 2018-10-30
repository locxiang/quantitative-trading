package okex_test

import (
	"testing"
	"github.com/locxiang/quantitative-trading/exchange/okex"
	"time"
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
)

func init() {
	Initialization.ServerInit()
}

func TestInit(t *testing.T) {
	okex.Init()

	time.Sleep(59 * time.Second)
}
