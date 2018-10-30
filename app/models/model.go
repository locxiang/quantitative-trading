package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/locxiang/quantitative-trading/app/util"
)

type BaseModel struct {
	Id        uint64 `gorm:"primary_key;unique_index;AUTO_INCREMENT;column:id" json:"id"`
	CreatedOn int64  `gorm:"column:created_on" json:"created_on"`
	UpdatedOn int64  `gorm:"column:updated_on" json:"updated_on"`
}

func (u *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", util.Time2Millisecond(time.Now()))
	scope.SetColumn("UpdatedOn", util.Time2Millisecond(time.Now()))
	return nil
}

func (u *BaseModel) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("UpdatedOn", util.Time2Millisecond(time.Now()))
	return nil
}
