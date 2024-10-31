package service

import (
	"ruoyi-go/app/model"
	"ruoyi-go/app/response"
	"ruoyi-go/framework/dal"
)

type SysUser struct{}

// 根据用户名获取用户详情
func (*SysUser) GetDetailByUserName(userName string) *response.SysUserDetail {

	var detail response.SysUserDetail

	dal.Gorm.Model(&model.SysUser{}).Where("user_name = ?", userName).Take(&detail)

	return &detail
}
