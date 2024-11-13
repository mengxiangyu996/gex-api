package service

import (
	"isme-go/app/model"
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/framework/dal"
)

type User struct{}

// 获取用户信息
func (*User) GetDetailByUsername(username string) response.User {

	var user response.User

	dal.Gorm.Model(&model.User{}).Where("username = ?", username).Take(&user)

	return user
}

// 获取用户信息
func (*User) GetDetailById(id int) response.User {

	var user response.User

	dal.Gorm.Model(&model.User{}).Where("id = ?", id).Take(&user)

	return user
}

// 获取用户列表
func (*User) Page(param request.UserPage) ([]response.User, int) {

	users := make([]response.User, 0)
	var count int64

	query := dal.Gorm.Model(&model.User{})

	if param.Username != "" {
		query = query.Where("username like ?", "%"+param.Username+"%")
	}

	if param.Enable != nil {
		query = query.Where("enable = ?", param.Enable)
	}

	if param.Gender != nil {
		query = query.Where("gender = ?", param.Gender)
	}

	query.Count(&count).Offset((param.PageNo - 1) * param.PageSize).Limit(param.PageSize).Scan(&users)

	return users, int(count)
}

// 删除用户
func (*User) Delete(id int) error {

	query := dal.Gorm.Begin()

	if err := query.Model(&model.User{}).Where("id = ?", id).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	if err := query.Model(&model.Profile{}).Where("user_id = ?", id).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	return query.Commit().Error
}

// 添加用户
func (*User) Insert(param request.UserAdd) error {

	user := model.User{
		Username: param.Username,
		Password: param.Password,
		Enable:   param.Enable,
	}

	query := dal.Gorm.Begin()

	if err := query.Model(&model.User{}).Create(&user).Error; err != nil {
		query.Rollback()
		return err
	}

	for _, roleId := range param.RoleIds {
		if err := query.Model(&model.UserRolesRole{}).Create(&model.UserRolesRole{
			UserId: user.Id,
			RoleId: roleId,
		}).Error; err != nil {
			query.Rollback()
			return err
		}
	}

	return query.Commit().Error
}

// 更新用户
func (*User) Update(user request.UserUpdate) error {

	query := dal.Gorm.Begin()

	if err := query.Model(&model.User{}).Select("enable").Where("id = ?", user.Id).Updates(&model.User{
		Password: user.Password,
		Enable:   user.Enable,
	}).Error; err != nil {
		query.Rollback()
		return err
	}

	if len(user.RoleIds) > 0 {
		if err := query.Model(&model.UserRolesRole{}).Where("user_id = ?", user.Id).Delete(nil).Error; err != nil {
			query.Rollback()
			return err
		}

		for _, roleId := range user.RoleIds {
			if err := query.Model(&model.UserRolesRole{}).Create(&model.UserRolesRole{
				UserId: user.Id,
				RoleId: roleId,
			}).Error; err != nil {
				query.Rollback()
				return err
			}
		}
	}

	return query.Commit().Error
}
