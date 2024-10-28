package router

import (
	"ruoyi-go/controller"

	"github.com/gin-gonic/gin"
)

func AdminApi(server *gin.Engine) {

	api := server.Group("")
	{
		api.GET("/captchaImage", (&controller.Common{}).CaptchaImage) // 生成验证码
	}

}
