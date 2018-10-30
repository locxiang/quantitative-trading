package binance

import (
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
	"testing"
	"time"
	"github.com/locxiang/quantitative-trading/exchange/binance"
)

func init() {
	Initialization.ServerInit()
}

func TestInit(t *testing.T) {
	done := binance.Init()

	time.Sleep(10 * time.Second)
	t.Log("结束")

	done <- struct{}{}
}
