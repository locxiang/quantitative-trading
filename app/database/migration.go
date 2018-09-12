package database

import (
	"github.com/locxiang/quantitative-trading/app/models/user"
	"github.com/locxiang/quantitative-trading/app/models/trade"
)

//数据库初始化
func MigrationData() {
	new(user.User).Migrate()

	t1 := &trade.Trade{
		Symbol: "EOSUSDT",
	}
	t1.Migrate()

	t2 := &trade.Trade{
		Symbol: "BTCUSDT",
	}
	t2.Migrate()
}
