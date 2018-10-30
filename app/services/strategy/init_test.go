package strategy

import (
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
	"testing"
	poolService "github.com/locxiang/quantitative-trading/app/services/pool"
)

func init() {
	Initialization.ServerInit()

	//TODO 启动数据池
	poolService.Init()
}

func TestInit2(t *testing.T) {
	Init()
}
