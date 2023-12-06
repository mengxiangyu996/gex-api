package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 访问用户数据
type User struct{}

// 创建用户
func (t *User) Create(user *model.User) error {
	return db.GormClient.Model(&model.User{}).Create(&user).Error
}
