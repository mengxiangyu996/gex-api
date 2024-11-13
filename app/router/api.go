package router

import (
	"isme-go/app/controller"
	"isme-go/app/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRegister(server *gin.Engine) {

	api := server.Group("/api")
	{
		api.POST("/auth/login", (&controller.Auth{}).Login)
		api.GET("/auth/captcha", (&controller.Auth{}).Captcha)
	}

	api = server.Group("/api", middleware.Authorization())
	{
		api.POST("/auth/current-role/switch/:roleCode", (&controller.Auth{}).SwitchCurrentRole)
		api.POST("/auth/logout", (&controller.Auth{}).Logout)
		api.GET("/user/detail", (&controller.User{}).Detail)
		api.GET("/user", (&controller.User{}).Page)
		api.DELETE("/user/:id", (&controller.User{}).Delete)
		api.GET("/role/permissions/tree", (&controller.Role{}).PermissionsTree)
	}

}
