package user

import (
	"testing"
	"github.com/go-ffmt/ffmt"
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
)

func init() {

	Initialization.ServerInit()
}

func TestGetUserAll(t *testing.T) {
	users := GetAll()

	ffmt.Puts(users)
}
