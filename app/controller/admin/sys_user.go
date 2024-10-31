package admin

import (
	"ruoyi-go/app/request"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/framework/message"
	"ruoyi-go/utils/captcha"
	"ruoyi-go/utils/password"

	"github.com/gin-gonic/gin"
)

type SysUser struct{}

// 登录
func (*SysUser) Login(ctx *gin.Context) {

	var param request.SysUserLogin

	if err := ctx.ShouldBind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	if !captcha.NewCaptcha().Verify(param.Uuid, param.Code) {
		message.Error(ctx, "验证码错误")
		return
	}

	if param.UserName == "" || param.Password == "" {
		message.Error(ctx, "用户名或密码错误")
		return
	}

	user := (&service.SysUser{}).GetDetailByUserName(param.UserName)
	if user.UserId <= 0 {
		message.Error(ctx, "用户不存在")
		return
	}

	if !password.Verify(param.Password, user.Password) {
		message.Error(ctx, "用户名或密码错误")
		return
	}

	token := token.GetClaims(user).GenerateToken()

	message.Success(ctx, map[string]interface{}{
		"token": token,
	})
}
