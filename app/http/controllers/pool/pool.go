package pool

import (
	"net/http"
	"github.com/locxiang/quantitative-trading/app/errors"
	"github.com/locxiang/quantitative-trading/app/http/controllers"
	"github.com/gin-gonic/gin"
	poolService "github.com/locxiang/quantitative-trading/app/services/pool"
)

// HealthCheck shows `OK` as the ping-pong result.
func GetAll(c *gin.Context) {

	data := poolService.GetTradePools()
	controllers.SendResponse(c, errors.OK, data, http.StatusOK)

}
