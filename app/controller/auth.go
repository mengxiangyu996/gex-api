package controller

import (
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/app/service"
	"isme-go/app/token"
	"isme-go/framework/message"
	"isme-go/utils"
	"isme-go/utils/captcha"
	"isme-go/utils/password"
	"strings"

	"github.com/gin-gonic/gin"
)

var captchaCache = make(map[string]string)

type Auth struct{}

// 验证码
func (*Auth) Captcha(ctx *gin.Context) {

	svg, code := captcha.New(&captcha.Config{
		Width:      80,
		Height:     40,
		FontSize:   20,
		CharsCount: 4,
	}).Generate()

	captchaCache[strings.ToLower(code)] = code

	// 返回svg文件
	ctx.Data(200, "image/svg+xml", svg)
}

// 登录
func (*Auth) Login(ctx *gin.Context) {

	var param request.Login

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	if param.Username == "" || param.Password == "" {
		message.Error(ctx, "用户名或密码不能为空")
		return
	}

	defer delete(captchaCache, strings.ToLower(param.Captcha))
	if _, ok := captchaCache[strings.ToLower(param.Captcha)]; !ok {
		message.Error(ctx, 10003, "验证码错误")
		return
	}

	user := (&service.User{}).GetDetailByUsername(param.Username)

	if !password.Verify(user.Password, param.Password) {
		message.Error(ctx, "用户名或密码错误")
		return
	}

	roleIds := (&service.UserRolesRole{}).GetRoleIdsByUserId(user.Id)
	roles := (&service.Role{}).GetListByIds(roleIds, true)
	if len(roles) <= 0 {
		message.Error(ctx, "用户未关联角色")
		return
	}

	roleCodes := make([]string, 0)
	for _, role := range roles {
		roleCodes = append(roleCodes, role.Code)
	}

	accessToken := token.GetClaims(response.UserToken{
		Id:              user.Id,
		Username:        user.Username,
		RoleCodes:       roleCodes,
		CurrentRoleCode: roleCodes[0],
	}).GenerateToken()

	message.Success(ctx, map[string]interface{}{
		"data": map[string]interface{}{
			"accessToken": accessToken,
		},
	})
}

// 切换当前角色
func (*Auth) SwitchCurrentRole(ctx *gin.Context) {

	roleCode := ctx.Param("roleCode")

	userId := ctx.GetInt("userId")
	username := ctx.GetString("username")

	roleIds := (&service.UserRolesRole{}).GetRoleIdsByUserId(userId)
	roles := (&service.Role{}).GetListByIds(roleIds, true)

	roleCodes := make([]string, 0)
	for _, role := range roles {
		roleCodes = append(roleCodes, role.Code)
	}

	if !utils.Contains(roleCodes, roleCode) {
		message.Error(ctx, "您目前暂无此角色，请联系管理员申请权限")
		return
	}

	accessToken := token.GetClaims(response.UserToken{
		Id:              userId,
		Username:        username,
		RoleCodes:       roleCodes,
		CurrentRoleCode: roleCode,
	}).GenerateToken()

	message.Success(ctx, map[string]interface{}{
		"data": map[string]interface{}{
			"accessToken": accessToken,
		},
	})
}

// 退出登录
func (*Auth) Logout(ctx *gin.Context) {
	message.Success(ctx, map[string]interface{}{
		"data": true,
	})
}

// 修改密码
func (*Auth) Password(ctx *gin.Context) {

	var param request.Password

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	if param.OldPassword == "" || param.NewPassword == "" {
		message.Error(ctx, "旧密码或新密码不能为空")
		return
	}

	user := (&service.User{}).GetDetailById(ctx.GetInt("userId"))
	if !password.Verify(user.Password, param.OldPassword) {
		message.Error(ctx, "旧密码错误")
		return
	}

	user.Password = password.Generate(param.NewPassword)

	if err := (&service.User{}).Update(request.UserUpdate{
		Id:       user.Id,
		Password: user.Password,
	}); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}
