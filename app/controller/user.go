package controller

import (
	"fmt"
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/app/service"
	"isme-go/framework/message"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{}

// 获取当前登录用户的详情信息
func (*User) Detail(ctx *gin.Context) {

	userId := ctx.GetInt("userId")

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
		userResponse.CurrentRole = roles[0]
	}

	message.Success(ctx, map[string]interface{}{
		"data": userResponse,
	})
}

// 获取用户列表
func (*User) Page(ctx *gin.Context) {

	fmt.Println(ctx.Query("pageSize"))

	var param request.UserPage

	if err := ctx.BindQuery(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	fmt.Println(param)

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

	message.Success(ctx, nil)
}
