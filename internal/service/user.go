package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 访问用户数据
type User struct{}

// 创建用户
func (t *User) CreateUser(user *model.User) error {

	err := db.GormClient.Model(&model.User{}).Create(&user).Error

	return err
}

// 更新用户
func (t *User) UpadteUser(user *model.User) error {

	err := db.GormClient.Model(&model.User{}).Where("id = ?", user.Id).Updates(&user).Error

	return err
}

// 删除用户
func (t *User) DeleteUser(id int) error {

	err := db.GormClient.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error

	return err
}

// 获取用户列表
// isAdmin 为-1时，查询全部
func (t *User) GetUserListByPage(page, size, isAdmin int) ([]*model.User, int) {

	var (
		list  []*model.User
		count int64
	)

	query := db.GormClient.Model(&model.User{}).Order("id desc")

	if isAdmin > -1 {
		query.Where("is_admin = ?", isAdmin)
	}

	query.Count(&count)

	query.Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 获取用户详情
func (t *User) GetUserById(id int) *model.User {

	var detail *model.User

	db.GormClient.Model(&model.User{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 获取用户详情
func (t *User) GetUserByUsername(username string) *model.User {

	var detail *model.User

	db.GormClient.Model(&model.User{}).Where("username = ?", username).Take(&detail)

	return detail
}
