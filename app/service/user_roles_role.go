package service

import (
	"isme-go/app/model"
	"isme-go/app/request"
	"isme-go/framework/dal"
	"isme-go/utils"
)

type UserRolesRole struct{}

// 获取用户角色列表
func (*UserRolesRole) GetRoleIdsByUserId(userId int) []int {

	ids := make([]int, 0)

	dal.Gorm.Model(&model.UserRolesRole{}).Where("user_id = ?", userId).Pluck("role_id", &ids)

	return ids
}

// 添加用户角色
func (u *UserRolesRole) Insert(param request.RoleUsersAdd) error {

	query := dal.Gorm.Begin()

	for _, userId := range param.UserIds {
		roleIds := u.GetRoleIdsByUserId(userId)
		if utils.Contains(roleIds, param.RoleId) {
			continue
		}
		if err := query.Create(&model.UserRolesRole{
			UserId: userId,
			RoleId: param.RoleId,
		}).Error; err != nil {
			query.Rollback()
			return err
		}
	}

	return query.Commit().Error
}

// 取消分配角色-批量
func (*UserRolesRole) Delete(param request.RoleUsersRemove) error {
	return dal.Gorm.Model(&model.UserRolesRole{}).Where("user_id in ? AND role_id = ?", param.UserIds, param.RoleId).Delete(nil).Error
}
