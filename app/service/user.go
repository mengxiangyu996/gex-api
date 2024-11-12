package service

import (
	"isme-go/app/model"
	"isme-go/framework/dal"
)

type User struct{}

// 获取用户信息
func (*User) GetDetailByUsername(username string) model.User {

	var user model.User

	dal.Gorm.Model(&model.User{}).Where("username = ?", username).Take(&user)

	return user
}
