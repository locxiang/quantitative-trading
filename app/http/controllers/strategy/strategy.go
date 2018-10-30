package strategy

import (
	"github.com/gin-gonic/gin"
	strategyService "github.com/locxiang/quantitative-trading/app/services/strategy"
	"net/http"
	"github.com/locxiang/quantitative-trading/app/http/controllers"
	"github.com/locxiang/quantitative-trading/app/errors"
	"github.com/locxiang/quantitative-trading/app/services/order"
)

// HealthCheck shows `OK` as the ping-pong result.
func GetAll(c *gin.Context) {

	data := strategyService.GetAll()
	controllers.SendResponse(c, errors.OK, data, http.StatusOK)

}

//返回利润
func ProfitList(c *gin.Context)  {
	data := order.ProfitList()
	controllers.SendResponse(c, errors.OK, data, http.StatusOK)
}