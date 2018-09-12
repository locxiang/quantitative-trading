package user

import (
	"github.com/locxiang/quantitative-trading/app/models/user"
	"github.com/locxiang/quantitative-trading/app/models"
)

//获取所有用户
func GetAll() (users []user.User) {


	models.DB.Find(&users)



	return users
}
