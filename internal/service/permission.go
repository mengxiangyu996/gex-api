package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 访问权限数据
type Permission struct{}

// 创建权限
func (t *Permission) CreatePermission(permission *model.Permission) error {

	err := db.GormClient.Model(&model.Permission{}).Create(&permission).Error

	return err
}

// 更新权限
func (t *Permission) UpdatePermission(permission *model.Permission) error {

	err := db.GormClient.Model(&model.Permission{}).Where("id = ?", permission.Id).Updates(&permission).Error

	return err
}

// 删除权限
func (t *Permission) DeletePermission(id int) error {

	err := db.GormClient.Model(&model.Permission{}).Where("id = ?", id).Delete(&model.Permission{}).Error

	return err
}

// 获取权限列表
func (t *Permission) GetPermissionListByPage(page, size int) ([]*model.Permission, int) {

	var (
		list []*model.Permission
		count int64
	)

	query := db.GormClient.Model(&model.Permission{}).Order("id desc")

	query.Count(&count)

	query.Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 获取权限列表
func (t *Permission) GetPermissionList() []*model.Permission {

	var list []*model.Permission

	db.GormClient.Model(&model.Permission{}).Order("id desc").Find(&list)

	return list
}

// 获取权限详情
func (t *Permission) GetPermissionById(id int) *model.Permission {

	var detail *model.Permission

	db.GormClient.Model(&model.Permission{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 获取权限详情
func (t *Permission) GetPermissionByPathWithMethod(path, method string) *model.Permission {

	var detail *model.Permission

	db.GormClient.Model(&model.Permission{}).Where("path = ?", path).Where("method = ?", method).Take(&detail)

	return detail
}