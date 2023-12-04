package service

import (
	"breeze-api/pkg/dal"
	"breeze-api/internal/model"

	"gorm.io/gorm"
)

// 访问用户数据
type User struct{}

// 创建用户
func (t *User) Create(user *model.User) error {
	return dal.DBClient.Model(&model.User{}).Create(&user).Error
}

// 用户详情
func (t *User) Detail(scopes ...func(*gorm.DB) *gorm.DB) *model.User {

	var user *model.User

	dal.DBClient.Model(&model.User{}).Scopes(scopes...).First(&user)

	return user
}
