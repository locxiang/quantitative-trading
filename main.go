package main

import (
	"github.com/locxiang/quantitative-trading/app/pkg/setting"
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/lexkong/log"
	"github.com/locxiang/quantitative-trading/app/database"
	"github.com/spf13/pflag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/locxiang/quantitative-trading/app/routers"
	"github.com/locxiang/quantitative-trading/app/http/middleware"
	"net/http"
)

var (
	cfg = pflag.StringP("config", "c", "config/app.ini", "app.ini config file path.")
)

func main() {
	var err error

	log.InitWithFile("config/log.yaml")

	//配置初始化
	err = setting.Init(*cfg)
	if err != nil {
		log.Fatal("配置加载失败: ", err)
	}
	fmt.Printf("配置加载：%#v \n", setting.Env())

	//数据库连接
	err = models.Init()
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	defer models.CloseDB()

	//TODO 数据库初始化
	database.MigrationData()



	//启动http服务
	gin.SetMode(setting.Env().RunMode)
	// Create the Gin engine.
	g := gin.New()
	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlewares.
		//middleware.Logging(),
		middleware.RequestId(),
	)

	add := setting.Env().Server.HttpListen
	err = http.ListenAndServe(add, g)
	log.Error("http服务出错:", err)

}
