package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/locxiang/quantitative-trading/app/util"
)

type BaseModel struct {
	Id        uint64 `gorm:"primary_key;unique_index;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"`
}

func (u *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", util.Time2Millisecond(time.Now()))
	scope.SetColumn("UpdatedAt", util.Time2Millisecond(time.Now()))
	return nil
}
