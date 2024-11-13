package service

import (
	"isme-go/app/model"
	"isme-go/app/response"
	"isme-go/framework/dal"
)

type Role struct{}

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
