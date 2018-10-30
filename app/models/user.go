package models

import (
	"github.com/jinzhu/gorm"
	"github.com/locxiang/quantitative-trading/app/util"
	"github.com/lexkong/log"
)

type User struct {
	BaseModel
	UserName string `gorm:"column:username;not null;unique_index;type:varchar(50);" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

func (u *User) TableName() string {
	return "users"
}

//数据迁移
func (u *User) Migrate() {

	exist := DB.HasTable(u)
	if exist {
		return
	}
	log.Infof("创建数据表：%s", u.TableName())

	//建表
	DB.AutoMigrate(u)

	//TODO  导入数据

	user1 := &User{
		UserName: "admin",
		Password: "123123",
	}

	DB.Create(user1)

}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	u.BaseModel.BeforeCreate(scope)
	scope.SetColumn("Password", util.PasswordEncrypt(u.Password))
	return nil
}
