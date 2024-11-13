package service

import (
	"isme-go/app/model"
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/framework/dal"
)

type Role struct{}

// 获取角色详情
func (*Role) GetDetailById(id int) response.Role {

	var role response.Role

	dal.Gorm.Model(&model.Role{}).Where("id = ?", id).Take(&role)

	return role
}

// 获取角色列表
func (*Role) GetListByIds(ids []int, enable bool) []response.Role {

	roles := make([]response.Role, 0)

	query := dal.Gorm.Model(&model.Role{}).Where("id in ?", ids)

	if enable {
		query = query.Where("enable = ?", 1)
	}

	query.Scan(&roles)

	return roles
}

// 根据编码获取角色
func (*Role) GetDetailByCode(code string) response.Role {

	var role response.Role

	dal.Gorm.Model(&model.Role{}).Where("code = ?", code).Scan(&role)

	return role
}

// 角色分页
func (*Role) Page(param request.RolePage) ([]response.Role, int) {

	roles := make([]response.Role, 0)
	var total int64

	query := dal.Gorm.Model(&model.Role{})

	if param.Name != "" {
		query = query.Where("name like ?", "%"+param.Name+"%")
	}

	if param.Enable != nil {
		query = query.Where("enable = ?", param.Enable)
	}

	query.Count(&total).Offset((param.PageNo - 1) * param.PageSize).Limit(param.PageSize).Scan(&roles)

	return roles, int(total)
}

// 获取角色列表
func (*Role) List() []response.Role {

	roles := make([]response.Role, 0)

	dal.Gorm.Model(&model.Role{}).Scan(&roles)

	return roles
}

// 添加角色
func (*Role) Insert(param request.RoleAdd) error {

	role := model.Role{
		Code:   param.Code,
		Name:   param.Name,
		Enable: param.Enable,
	}

	query := dal.Gorm.Begin()

	if err := query.Model(&model.Role{}).Select("enable").Create(&role).Error; err != nil {
		query.Rollback()
		return err
	}

	for _, permissionId := range param.PermissionIds {
		if err := query.Model(&model.RolePermissionsPermission{}).Create(&model.RolePermissionsPermission{
			RoleId:       role.Id,
			PermissionId: permissionId,
		}).Error; err != nil {
			query.Rollback()
			return err
		}
	}

	return query.Commit().Error
}

// 修改角色
func (*Role) Update(param request.RoleUpdate) error {

	role := model.Role{
		Id:     param.Id,
		Name:   param.Name,
		Enable: param.Enable,
	}

	query := dal.Gorm.Begin()

	if err := query.Model(&model.Role{}).Where("id = ?", role.Id).Updates(&role).Error; err != nil {
		query.Rollback()
		return err
	}

	if err := query.Model(&model.RolePermissionsPermission{}).Where("role_id = ?", role.Id).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	for _, permissionId := range param.PermissionIds {
		if err := query.Model(&model.RolePermissionsPermission{}).Create(&model.RolePermissionsPermission{
			RoleId:       role.Id,
			PermissionId: permissionId,
		}).Error; err != nil {
			query.Rollback()
			return err
		}
	}

	return query.Commit().Error
}

// 删除角色
func (*Role) Delete(id int) error {

	query := dal.Gorm.Begin()

	if err := query.Model(&model.Role{}).Where("id = ?", id).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	if err := query.Model(&model.RolePermissionsPermission{}).Where("role_id = ?", id).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	return query.Commit().Error
}
