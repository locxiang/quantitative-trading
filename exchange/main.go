package main


import (
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
	"github.com/locxiang/quantitative-trading/exchange/huobi"
	"github.com/locxiang/quantitative-trading/exchange/okex"
)

func init() {
	Initialization.ServerInit()
}

func main() {
	go huobi.Init()

	go okex.Init()

	select {

	}
}
