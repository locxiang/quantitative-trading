package models

import (
	"time"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
}

//func (u *BaseModel) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("created_at", time.Now())
//	scope.SetColumn("updated_at", time.Now().AddDate(3,3,3))
//	return nil
//}
