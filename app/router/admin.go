package router

import (
	"ruoyi-go/app/controller"

	"github.com/gin-gonic/gin"
)

func AdminApi(server *gin.Engine) {

	api := server.Group("api")
	{
		api.GET("/common/captcha", (&controller.Common{}).Captcha) // 生成验证码
	}

}
