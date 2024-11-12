package controller

import (
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/app/service"
	"isme-go/app/token"
	"isme-go/framework/message"
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
		message.Error(ctx, "验证码错误")
		return
	}

	user := (&service.User{}).GetDetailByUsername(param.Username)

	if !password.Verify(param.Password, user.Password) {
		message.Error(ctx, "用户名或密码错误")
		return
	}

	token := token.GetClaims(response.UserToken{
		Id:       user.Id,
		Username: user.Username,
	}).GenerateToken()

	message.Success(ctx, "登录成功", map[string]interface{}{
		"accessToken": token,
		"originUrl":   strings.Replace(ctx.FullPath(), "/api", "", 1),
	})
}
