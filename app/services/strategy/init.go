package strategy

import (
	"github.com/locxiang/quantitative-trading/app/services/pool"
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/lexkong/log"
)

func Init() {

	strategies := GetAll()

	for _, p := range pool.GetTradePools() {
		for _, s := range strategies {
			ps := Convert(s)
			p.AddStrategy(ps)
		}

	}

}

func Convert(strategy *models.Strategy) (s pool.Strategy) {
	switch strategy.Type {
	case models.STRATEGY_TYPE_MERCHANDISER:
		m := new(Merchandiser)
		m.Id = strategy.Id
		m.BuyQuantity = strategy.BuyQuantity
		m.Ratio = strategy.Ratio
		m.GoodRatio = strategy.GoodRatio
		m.BadRatio = strategy.BadRatio

		s = m
	default:
		log.Errorf(nil, "没有找到对应策略")
	}

	return
}

func GetAll() (strategies []*models.Strategy) {
	models.DB.Find(&strategies)
	return
}
