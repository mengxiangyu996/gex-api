package service

import (
	"isme-go/app/dto"
	"isme-go/app/model"
	"isme-go/framework/dal"
)

type Profile struct{}

// 获取用户详情
func (*Profile) GetDetailByUserId(userId int) dto.ProfileResponse {

	var profile dto.ProfileResponse

	dal.Gorm.Model(&model.Profile{}).Where("user_id = ?", userId).Take(&profile)

	return profile
}

// 更新用户资料
func (*Profile) Update(param dto.UserProfileUpdateRequest) error {
	return dal.Gorm.Model(&model.Profile{}).Select("gender").Where("user_id = ?", param.Id).Updates(&param).Error
}
