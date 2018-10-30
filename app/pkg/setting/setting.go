package setting

import (
	"time"

	"github.com/go-ini/ini"
)

type EnvS struct {
	RunMode  string   `ini:"RUN_MODE"`
	Server   Server   `ini:"server"`
	Database Database `ini:"database"`
	InfluxDB InfluxDB `ini:"influxDB"`
	Exchange Exchange `ini:"exchange"`
}

type Exchange struct {
	Type      string
	Url       string
	ApiKey    string
	SecretKey string
	Spot      []string
}

type Server struct {
	HttpListen  string        `ini:"HTTP_LISTEN"`
	ReadTimeout time.Duration `ini:"READ_TIMEOUT"`
}

type Database struct {
	Type     string `ini:"TYPE"`
	User     string `ini:"USER"`
	Password string `ini:"PASSWORD"`
	Host     string `ini:"HOST"`
	Db       string `ini:"DB"`
}

type InfluxDB struct {
	User      string `ini:"USER"`
	Password  string `ini:"PASSWORD"`
	Addr      string `ini:"ADDR"`
	Db        string `ini:"DB"`
	Precision string `ini:"Precision"`
}

var (
	e EnvS
)

//避免全局配置被修改
func Env() EnvS {
	return e
}

//加载配置文件
func Init(file string) error {
	var err error

	p := new(EnvS)
	err = ini.MapTo(p, file)
	if err != nil {
		return err
	}

	e = *p
	return nil
}
