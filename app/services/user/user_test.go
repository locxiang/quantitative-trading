package user

import (
	"testing"
	"github.com/locxiang/quantitative-trading/app/pkg/setting"
	"github.com/lexkong/log"
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/go-ffmt/ffmt"
)

func init() {

	var err error
	err = log.InitWithFile("config/log.yaml")
	err = setting.Init("config/app.ini")
	err = models.Init()
	if err != nil {
		log.Fatal("", err)
	}
}

func TestGetUserAll(t *testing.T) {
	users := GetAll()

	ffmt.Puts(users)
}
