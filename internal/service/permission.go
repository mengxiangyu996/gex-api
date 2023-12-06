package service

import (
	"breeze-api/internal/model"
	"breeze-api/pkg/db"
)

// 访问权限数据
type permission struct{}

// 创建权限
func (t *permission) Createpermission(permission *model.Permission) error {

	err := db.GormClient.Model(&model.Permission{}).Create(&permission).Error

	return err
}

// 更新权限
func (t *permission) Updatepermission(permission *model.Permission) error {

	err := db.GormClient.Model(&model.Permission{}).Where("id = ?", permission.Id).Updates(&permission).Error

	return err
}

// 删除权限
func (t *permission) Deletepermission(id int) error {

	err := db.GormClient.Model(&model.Permission{}).Where("id = ?", id).Delete(&model.Permission{}).Error

	return err
}

// 获取权限列表
func (t *permission) GetpermissionListByPage(page, size int) ([]*model.Permission, int) {

	var (
		list []*model.Permission
		count int64
	)

	query := db.GormClient.Model(&model.Permission{})

	query.Count(&count)

	query.Limit(size).Offset((page - 1) * size).Find(&list)

	return list, int(count)
}

// 获取权限详情
func (t *permission) GetpermissionById(id int) *model.Permission {

	var detail *model.Permission

	db.GormClient.Model(&model.Permission{}).Where("id = ?", id).Take(&detail)

	return detail
}