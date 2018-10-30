package models

import (
	"github.com/lexkong/log"
	"github.com/locxiang/quantitative-trading/app/util"
)

/**
TODO 计算盈利 sql

 */
type Strategy struct {
	BaseModel
	Type string `json:"type"` //策略类型
	*StrategyMerchandiser
}

const (
	STRATEGY_TYPE_MERCHANDISER = "StrategyMerchandiser"
)

//跟单策略
type StrategyMerchandiser struct {
	Ratio       float64 `json:"ratio"`        //涨幅比例
	BuyQuantity float64 `json:"buy_quantity"` //购买数量
	GoodRatio   float64 `json:"good_ratio"`   //止盈比例
	BadRatio    float64 `json:"bad_ratio"`    //止损比例
}

func (s *Strategy) TableName() string {
	return "strategies"
}

func (s *Strategy) Migrate() {
	exist := DB.HasTable(s)
	if exist {
		return
	}
	log.Infof("创建数据表：%s", s.TableName())

	//建表
	DB.AutoMigrate(s)

	//创建策略1 Merchandiser
	for i := 0; i < 100; i++ {
		data := new(Strategy)
		data.StrategyMerchandiser = new(StrategyMerchandiser)
		data.StrategyMerchandiser.randData()
		data.Type = data.StrategyMerchandiser.Type()
		DB.FirstOrCreate(data, data)
	}
}

func (m *StrategyMerchandiser) Type() string {
	return STRATEGY_TYPE_MERCHANDISER
}

//随机生成数据
func (m *StrategyMerchandiser) randData() {
	var r int
	r = util.RandInt(1, 300)
	m.Ratio = float64(r) / 10000

	r = util.RandInt(1, 5)
	m.GoodRatio = float64(r) / 10000

	r = util.RandInt(1, 5)
	m.BadRatio = float64(r) / 10000 * -1

	m.BuyQuantity = 10
}
