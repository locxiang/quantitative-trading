package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/locxiang/quantitative-trading/app/pkg/setting"
	"time"
	"github.com/lexkong/log"
)

var DB *gorm.DB
var dbConnError = make(chan error)

func Init() error {
	err := Conn()
	if err != nil {
		return err
	}

	go checkServer()
	return nil
}

func Conn() error {
	var err error

	conf := setting.Env().Database

	dbStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Db)

	DB, err = gorm.Open(conf.Type, dbStr)

	if err != nil {
		return err
	}

	log.Infof("数据库连接成功：%s", conf.Host)

	t := setting.Env().RunMode == "debug"
	DB.LogMode(t)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	return nil
}

func CloseDB() {
	defer DB.Close()
}

//实时检查数据库连接情况
func checkServer() {
	for {
		select {
		case err := <-dbConnError:
			//TODO gorm 自带重连机制，是否要处理什么
			log.Error("数据库连接异常", err)
			//Conn()

		case <-time.After(1 * time.Second):
			go pingServer()
		}
	}
}

func pingServer() {
	err := DB.DB().Ping()
	if err != nil {
		dbConnError <- err
	}
}
