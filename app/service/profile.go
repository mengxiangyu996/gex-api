package service

import (
	"isme-go/app/model"
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/framework/dal"
)

type Profile struct{}

// 获取用户详情
func (*Profile) GetDetailByUserId(userId int) response.Profile {

	var profile response.Profile

	dal.Gorm.Model(&model.Profile{}).Where("user_id = ?", userId).Take(&profile)

	return profile
}

// 更新用户资料
func (*Profile) Update(param request.UserProfileUpdate) error {
	return dal.Gorm.Model(&model.Profile{}).Select("gender").Where("user_id = ?", param.Id).Updates(&param).Error
}
