package database

import (
	"github.com/locxiang/quantitative-trading/app/pkg/Initialization"
	"testing"
)

func init()  {
	Initialization.ServerInit()
}

func TestMigrationData(t *testing.T) {
	MigrationData()
}