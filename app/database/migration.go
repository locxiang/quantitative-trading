package database

import (
	"github.com/locxiang/quantitative-trading/app/models"
)

//数据库初始化
func MigrationData() {
	new(models.User).Migrate()
	new(models.Pool).Migrate()
	new(models.Order).Migrate()
	new(models.Strategy).Migrate()

	t1 := &models.Trade{
		Symbol: "EOSUSDT",
	}
	t1.Migrate()

	t2 := &models.Trade{
		Symbol: "BTCUSDT",
	}
	t2.Migrate()
}
