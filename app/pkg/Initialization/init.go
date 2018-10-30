package Initialization

import (
	"github.com/lexkong/log"
	"github.com/locxiang/quantitative-trading/app/pkg/setting"
	"github.com/locxiang/quantitative-trading/app/models"

	"github.com/locxiang/quantitative-trading/app/influxdb"
)

func ServerInit() {

	if err := log.InitWithFile("config/log.yaml"); err != nil {
		log.Fatal("", err)
	}

	if err := setting.Init("config/app.ini"); err != nil {
		log.Fatal("", err)
	}

	if err := models.Init(); err != nil {
		log.Fatal("", err)
	}


	if err := influxdb.Init(); err != nil {
		log.Fatal("", err)
	}

}
