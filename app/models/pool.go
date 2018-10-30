package models

import (
	"time"
	"github.com/lexkong/log"
)

type Pool struct {
	BaseModel
	Duration time.Duration `json:"duration"` //在一定时间内
	Symbol   string        `json:"symbol"`   //类别
}

func (p *Pool) TableName() string {
	return "pools"
}

func (p *Pool) Migrate() {
	exist := DB.HasTable(p)
	if exist {
		return
	}
	log.Infof("创建数据表：%s", p.TableName())

	//建表
	DB.AutoMigrate(p)

	//生产数据
	p.crateData()
}

func (p *Pool) crateData() {
	for i := 3; i < 10; i++ {
		data := new(Pool)
		data.Duration = time.Duration(i) * time.Second
		data.Symbol = "EOSUSDT"
		DB.Create(data)
	}
}
