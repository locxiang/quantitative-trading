package pool

import (
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
	"testing"
	"time"
)

func init() {
	Initialization.ServerInit()
	Init()
}

func TestLoadDB(t *testing.T) {
	LoadStartDB()
}

func TestAdd(t *testing.T) {
	Add("EOSUSDT", 8*time.Second)
	Add("EOSUSDT", 7*time.Second)
	Add("EOSUSDT", 8*time.Second)
}
