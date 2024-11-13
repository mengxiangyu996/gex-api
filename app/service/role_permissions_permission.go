package service

import (
	"isme-go/app/model"
	"isme-go/framework/dal"
)

type RolePermissionsPermission struct{}

// 根据角色id获取资源id
func (*RolePermissionsPermission) GetPermissionIdsByRoleIds(roleIds []int) []int {

	ids := make([]int, 0)

	dal.Gorm.Model(&model.RolePermissionsPermission{}).Where("role_id in ?", roleIds).Pluck("permission_id", &ids)

	return ids
}

// 检查角色是否拥有权限
func (*RolePermissionsPermission) CheckHasPermission(roleId, permissionId int) bool {

	var count int64

	dal.Gorm.Model(&model.RolePermissionsPermission{}).Where("role_id = ? and permission_id = ?", roleId, permissionId).Count(&count)

	return count > 0
}
