package router

import (
	"ruoyi-go/app/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminApi(server *gin.Engine) {

	api := server.Group("admin/api")
	{
		api.GET("/captcha/image", (&admin.Captcha{}).Image) // 生成验证码
		api.POST("/user/login", (&admin.SysUser{}).Login)   // 登录
	}

}
