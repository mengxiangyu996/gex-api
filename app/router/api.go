package router

import (
	"isme-go/app/controller"

	"github.com/gin-gonic/gin"
)

func ApiRegister(server *gin.Engine) {

	api := server.Group("/api")
	{
		api.GET("/auth/captcha", (&controller.Auth{}).Captcha)
		api.POST("/auth/login", (&controller.Auth{}).Login)
	}

}
