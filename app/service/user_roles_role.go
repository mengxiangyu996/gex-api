package service

import (
	"isme-go/app/model"
	"isme-go/framework/dal"
)

type UserRolesRole struct{}

// 获取用户角色列表
func (*UserRolesRole) GetRoleIdsByUserId(userId int) []int {

	ids := make([]int, 0)

	dal.Gorm.Model(&model.UserRolesRole{}).Where("user_id = ?", userId).Pluck("role_id", &ids)

	return ids
}
