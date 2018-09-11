package database

import "github.com/locxiang/quantitative-trading/app/models/user"

//数据库初始化
func MigrationData() {
	new(user.User).Migrate()
}


