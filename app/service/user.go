package service

import (
	"gex-api/app/model"
	"gex-api/app/request"
	"gex-api/app/response"
	"gex-api/pkg/dal"
)

// 用户
type User struct{}

// 创建用户
func (*User) Create(user *request.CreateUser) int {

	data := &model.User{
		Role:     user.Role,
		Username: user.Username,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		Avatar:   user.Avatar,
		Status:   user.Status,
	}

	if err := dal.Gorm.Model(&model.User{}).Create(&data).Error; err != nil {
		return 0
	}

	return data.Id
}

// 更新用户
func (*User) Update(user *request.UpdateUser) int {

	data := &model.User{
		Role:     user.Role,
		Username: user.Username,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		Avatar:   user.Avatar,
		Status:   user.Status,
	}

	if err := dal.Gorm.Model(&model.User{}).Where("id = ?", user.Id).Updates(&data).Error; err != nil {
		return 0
	}

	return user.Id
}

// 删除用户
func (*User) DeleteById(id int) error {
	return dal.Gorm.Model(&model.User{}).Where("id = ?", id).Delete(nil).Error
}

// 用户列表
func (*User) GetList(param *request.QueryListUser) ([]*response.UserList, int) {

	var count int64
	list := make([]*response.UserList, 0)

	query := dal.Gorm.Model(&model.User{})

	if param.Role > 0 {
		query.Where("role = ?", param.Role)
	}

	if param.Username != "" {
		query.Where("username like ?", "%"+param.Username+"%")
	}

	if param.Nickname != "" {
		query.Where("nickname like ?", "%"+param.Nickname+"%")
	}

	if param.Phone != "" {
		query.Where("phone like ?", "%"+param.Phone+"%")
	}

	if param.Email != "" {
		query.Where("email like ?", "%"+param.Email+"%")
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

// 用户详情
func (*User) GetDetailById(id int) *response.UserDetail {

	var detail *response.UserDetail

	dal.Gorm.Model(&model.User{}).Omit("password").Where("id = ?", id).Take(&detail)

	return detail
}

// 用户详情
func (*User) GetDetailByUsername(username string) *response.UserDetail {

	var detail *response.UserDetail

	dal.Gorm.Model(&model.User{}).Where("username = ?", username).Take(&detail)

	return detail
}

// 用户绑定角色
func (*User) BindRole(param *request.AdminBindRole) error {

	query := dal.Gorm.Begin()

	if err := query.Model(&model.UserRole{}).Where("user_id = ?", param.UserId).Delete(nil).Error; err != nil {
		query.Rollback()
		return err
	}

	for _, roleId := range param.RoleIds {
		if err := query.Model(&model.UserRole{}).Create(&model.UserRole{
			UserId: param.UserId,
			RoleId: roleId,
		}).Error; err != nil {
			query.Rollback()
			return err
		}
	}

	return query.Commit().Error
}

// 用户角色列表
func (*User) GetBindRole(userId int) []int {

	var roleIds []int

	dal.Gorm.Model(&model.UserRole{}).Where("user_id = ?", userId).Pluck("role_id", &roleIds)

	return roleIds
}
