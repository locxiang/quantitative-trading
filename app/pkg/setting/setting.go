package setting

import (
	"time"

	"github.com/go-ini/ini"
)

type env struct {
	RunMode  string   `ini:"RUN_MODE"`
	Server   Server   `ini:"server"`
	Database Database `ini:"database"`
}

type Server struct {
	HttpListen    string           `ini:"HTTP_LISTEN"`
	ReadTimeout time.Duration `ini:"READ_TIMEOUT"`
}

type Database struct {
	Type     string `ini:"TYPE"`
	User     string `ini:"USER"`
	Password string `ini:"PASSWORD"`
	Host     string `ini:"HOST"`
	Db       string `ini:"DB"`
}

var (
	e   env
)

//避免全局配置被修改
func Env() env {
	return e
}

//加载配置文件
func Init(file string) error {
	var err error

	p := new(env)
	err = ini.MapTo(p, file)
	if err != nil {
		return err
	}

	e = *p
	return nil
}
