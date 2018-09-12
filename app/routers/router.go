package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/locxiang/quantitative-trading/app/http/controllers"
	"github.com/locxiang/quantitative-trading/app/http/controllers/check"
	"net/http"
	"github.com/locxiang/quantitative-trading/app/http/middleware"
	"github.com/locxiang/quantitative-trading/app/http/controllers/user"
	"github.com/locxiang/quantitative-trading/app/errors"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		controllers.SendResponse(c, errors.RouterNot, nil, http.StatusNotFound)
		return
	})

	// pprof router
	pprof.Register(g)

	//不带认证中间件
	route := g.Group("/api/")
	{
		//检查系统服务 api
		route.GET("sd/health", check.HealthCheck)

		route.GET("users", user.GetAll)

	}

	////带认证中间件
	//api := g.Group("/v1/", middleware.AuthMiddleware())
	//{
	//	//user.
	//	api.GET("user", user.Info)
	//
	//}

	////websocket
	//websocket := g.Group("/v1/")
	//{
	//	websocket.GET("task/websocket/:task", task.Websocket)
	//}

	return g
}
