package trade

import (
	"github.com/locxiang/quantitative-trading/app/models"
)

func Add(t *models.Trade) {
	models.DB.Create(t)
}
