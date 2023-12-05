package service

import (
	"breeze-api/pkg/db"
	"breeze-api/internal/model"

	"gorm.io/gorm"
)

// 访问用户数据
type User struct{}

// 创建用户
func (t *User) Create(user *model.User) error {
	return db.GormClient.Model(&model.User{}).Create(&user).Error
}

// 用户详情
func (t *User) Detail(scopes ...func(*gorm.DB) *gorm.DB) *model.User {

	var user *model.User

	db.GormClient.Model(&model.User{}).Scopes(scopes...).First(&user)

	return user
}
