package controller

import (
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/app/service"
	"isme-go/framework/message"
	"isme-go/utils/password"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{}

// 获取当前登录用户的详情信息
func (*User) Detail(ctx *gin.Context) {

	userId := ctx.GetInt("userId")
	roleCode := ctx.GetString("roleCode")

	user := (&service.User{}).GetDetailById(userId)

	userResponse := response.UserDetail{
		User: user,
	}

	roleIds := (&service.UserRolesRole{}).GetRoleIdsByUserId(user.Id)
	roles := (&service.Role{}).GetListByIds(roleIds, true)

	userResponse.Roles = roles

	profile := (&service.Profile{}).GetDetailByUserId(user.Id)
	userResponse.Profile = profile

	if len(roles) > 0 {
		for _, role := range roles {
			if role.Code == roleCode {
				userResponse.CurrentRole = role
				break
			}
		}
	}

	message.Success(ctx, map[string]interface{}{
		"data": userResponse,
	})
}

// 获取用户列表
func (*User) Page(ctx *gin.Context) {

	var param request.UserPage

	if err := ctx.BindQuery(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	userPages := make([]response.UserPage, 0)

	users, count := (&service.User{}).Page(param)
	if len(users) > 0 {
		for _, user := range users {
			var userPage response.UserPage
			userPage.User = user
			profile := (&service.Profile{}).GetDetailByUserId(user.Id)
			userPage.Gender = profile.Gender
			userPage.Avatar = profile.Avatar
			userPage.Address = profile.Address
			userPage.Email = profile.Email
			roleIds := (&service.UserRolesRole{}).GetRoleIdsByUserId(user.Id)
			roles := (&service.Role{}).GetListByIds(roleIds, true)
			userPage.Roles = roles
			userPages = append(userPages, userPage)
		}
	}

	message.Success(ctx, map[string]interface{}{
		"data": map[string]interface{}{
			"pageData": userPages,
			"total":    count,
		},
	})
}

// 删除用户
func (*User) Delete(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	if id == 1 {
		message.Error(ctx, "不能删除根用户")
		return
	}

	if id == ctx.GetInt("userId") {
		message.Error(ctx, "非法操作，不能删除自己")
		return
	}

	if err := (&service.User{}).Delete(id); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}

// 添加用户
func (*User) Add(ctx *gin.Context) {

	var param request.UserAdd

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	if param.Username == "" {
		message.Error(ctx, "用户名不能为空")
		return
	}

	if param.Password == "" {
		message.Error(ctx, "密码不能为空")
		return
	}

	user := (&service.User{}).GetDetailByUsername(param.Username)
	if user.Id > 0 {
		message.Error(ctx, "用户名已存在")
		return
	}

	param.Password = password.Generate(param.Password)

	if err := (&service.User{}).Insert(param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}

// 修改资料
func (*User) ProfileUpdate(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	var param request.UserProfileUpdate

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	userId := ctx.GetInt("userId")

	if id != userId {
		message.Error(ctx, "越权操作，用户资料只能本人修改")
		return
	}

	param.Id = id

	if err := (&service.Profile{}).Update(param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}

// 修改密码
func (*User) PasswordReset(ctx *gin.Context) {

	var param request.UserUpdate

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	userId, _ := strconv.Atoi(ctx.Param("id"))
	if userId <= 0 {
		userId = ctx.GetInt("userId")
	}

	if err := (&service.User{}).Update(request.UserUpdate{
		Id:       userId,
		Password: password.Generate(param.Password),
	}); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}

// 修改用户
func (*User) Update(ctx *gin.Context) {

	var param request.UserUpdate

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	param.Id, _ = strconv.Atoi(ctx.Param("id"))

	if err := (&service.User{}).Update(param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}
