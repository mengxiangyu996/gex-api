package service

import (
	"gex-api/app/model"
	"gex-api/app/request"
	"gex-api/app/response"
	"gex-api/pkg/dal"
)

// 角色
type Role struct{}

// 创建角色
func (*Role) Create(role *request.CreateRole) error {
	return dal.Gorm.Model(&model.Role{}).Create(&model.Role{
		Name:   role.Name,
		Status: role.Status,
	}).Error
}

// 更新角色
func (*Role) Update(role *request.UpdateRole) error {
	return dal.Gorm.Model(&model.Role{}).Where("id = ?", role.Id).Updates(&model.Role{
		Name:   role.Name,
		Status: role.Status,
	}).Error
}

// 删除角色
func (*Role) DeleteById(id int) error {
	return dal.Gorm.Model(&model.Role{}).Where("id = ?", id).Delete(nil).Error
}

// 角色列表
func (*Role) GetList(param *request.QueryListRole) ([]*response.RoleList, int) {

	var count int64
	list := make([]*response.RoleList, 0)

	query := dal.Gorm.Model(&model.Role{})

	if param.Name != "" {
		query.Where("name like ?", "%"+param.Name+"%")
	}

	if param.Status > 0 {
		query.Where("status = ?", param.Status)
	}

	if param.IsPaging {
		query.Count(&count).Limit(param.Size).Offset((param.Page - 1) * param.Size)
	}

	query.Scan(&list)

	return list, int(count)
}

// 角色详情
func (*Role) GetDetailById(id int) *response.RoleDetail {

	var detail *response.RoleDetail

	dal.Gorm.Model(&model.Role{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 角色选项
func (*Role) Option() []*response.Option {

	list := make([]*response.Option, 0)

	dal.Gorm.Model(&model.Role{}).Select("id as value", "name as label").Where("status = ?", 1).Scan(&list)

	return list
}

// 角色绑定菜单
func (*Role) BindMenu(param *request.RoleBindMenu) error {

	query := dal.Gorm.Begin()

	if err := query.Model(&model.RoleMenu{}).Where("role_id = ?", param.RoleId).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	for _, menuId := range param.MenuIds {
		if err := query.Model(&model.RoleMenu{}).Create(&model.RoleMenu{
			RoleId: param.RoleId,
			MenuId: menuId,
		}).Error; err != nil {
			query.Rollback()
			return err
		}
	}

	return query.Commit().Error
}

// 角色菜单列表
func (*Role) GetBindMenu(roleId int) []int {

	var menuIds []int

	dal.Gorm.Model(&model.RoleMenu{}).Where("role_id = ?", roleId).Pluck("menu_id", &menuIds)

	return menuIds
}

// 角色绑定权限
func (*Role) BindPermission(param *request.RoleBindPermission) error {

	query := dal.Gorm.Begin()

	if err := query.Model(&model.RolePermission{}).Where("role_id = ?", param.RoleId).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	for _, permissionId := range param.PermissionIds {
		if err := query.Model(&model.RolePermission{}).Create(&model.RolePermission{
			RoleId:       param.RoleId,
			PermissionId: permissionId,
		}).Error; err != nil {
			query.Rollback()
			return err
		}
	}

	return query.Commit().Error
}

// 角色权限列表
func (*Role) GetBindPermission(roleId int) []int {

	var permissionIds []int

	dal.Gorm.Model(&model.RolePermission{}).Where("role_id = ?", roleId).Pluck("permission_id", &permissionIds)

	return permissionIds
}
