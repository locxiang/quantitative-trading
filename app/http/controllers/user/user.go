package user

import (
	userService "github.com/locxiang/quantitative-trading/app/services/user"
	"net/http"
	"github.com/go-ffmt/ffmt"
	"github.com/locxiang/quantitative-trading/app/errors"
	"github.com/locxiang/quantitative-trading/app/http/controllers"
	"github.com/gin-gonic/gin"
)

// HealthCheck shows `OK` as the ping-pong result.
func GetAll(c *gin.Context) {

	users := userService.GetAll()

	ffmt.Pjson(users)

	controllers.SendResponse(c, errors.OK, users, http.StatusOK)

}
