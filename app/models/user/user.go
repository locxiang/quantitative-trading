package user

import (
	"github.com/locxiang/quantitative-trading/app/models"
	"github.com/lexkong/log"
	"github.com/jinzhu/gorm"
	"github.com/locxiang/quantitative-trading/app/util"
)

type User struct {
	models.BaseModel
	UserName string `gorm:"column:username" json:"-"`
	Password string `gorm:"column:password" json:"-"`
}

//数据迁移
func (u *User) Migrate() {

	exist := models.DB.HasTable(u)
	if exist {
		log.Debug("数据库存在不用再导入")
		return
	}

	//建表
	models.DB.AutoMigrate(u)

	//建索引
	models.DB.Model(u).AddUniqueIndex("idx_user_name", "username")

	//TODO  导入数据

	user1 := &User{
		UserName: "admin",
		Password: "123123",
	}

	models.DB.Create(user1)

}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	//u.BaseModel.BeforeCreate(scope)
	scope.SetColumn("password", util.PasswordEncrypt(u.Password))
	return nil
}
