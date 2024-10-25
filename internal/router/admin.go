package router

import (
	"ruoyi-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func AdminApi(server *gin.Engine) {

	api := server.Group("")
	{
		api.GET("/captchaImage", (&handler.Common{}).CaptchaImage) // 生成验证码
	}

}
