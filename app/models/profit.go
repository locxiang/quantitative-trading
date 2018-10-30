package models

import "github.com/lexkong/log"

type Profit struct {
	BaseModel
	StrategyId uint64  `json:"strategy_id"`
	PoolId     uint64  `json:"pool_id"`
	Profit     float64 `json:"profit"`
	BuyCount   uint64  `json:"buy_count"`
}

func (p *Profit) TableName() string {
	return "profits"
}

func (p *Profit) Migrate() {
	exist := DB.HasTable(p)
	if exist {
		return
	}
	log.Infof("创建数据表：%s", p.TableName())

	//建表
	DB.AutoMigrate(p)

}