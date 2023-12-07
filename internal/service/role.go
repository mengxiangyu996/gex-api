package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 访问角色数据
type Role struct{}

// 创建角色
func (t *Role) CreateRole(role *model.Role) error {

	err := db.GormClient.Model(&model.Role{}).Create(&role).Error

	return err
}

// 更新角色
func (t *Role) UpdateRole(role *model.Role) error {

	err := db.GormClient.Model(&model.Role{}).Where("id = ?", role.Id).Updates(&role).Error

	return err
}

// 删除角色
func (t *Role) DeleteRole(id int) error {

	err := db.GormClient.Model(&model.Role{}).Where("id = ?", id).Delete(&model.Role{}).Error

	return err
}

// 获取角色列表
func (t *Role) GetRoleListByPage(page, size int) ([]*model.Role, int) {

	var (
		list []*model.Role
		count int64
	)

	query := db.GormClient.Model(&model.Role{}).Order("id desc")

	query.Count(&count)

	query.Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 获取角色列表
func (t *Role) GetRoleListByIds(ids []int) []*model.Role {

	var list []*model.Role

	db.GormClient.Model(&model.Role{}).Order("id desc").Where("id in ?", ids).Find(&list)

	return list
}

// 获取角色详情
func (t *Role) GetRoleById(id int) *model.Role {

	var detail *model.Role

	db.GormClient.Model(&model.Role{}).Where("id = ?", id).Take(&detail)

	return detail
}