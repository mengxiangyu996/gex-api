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
		api.PATCH("/user/profile/:id", (&controller.User{}).ProfileUpdate)
		api.POST("/auth/password", (&controller.Auth{}).Password)
		api.POST("/user", (&controller.User{}).Add)
		api.GET("/role/permissions/tree", (&controller.Role{}).PermissionsTree)
		api.GET("/role/page", (&controller.Role{}).Page)
		api.GET("/role", (&controller.Role{}).List)
		api.POST("/role", (&controller.Role{}).Add)
		api.PATCH("/role/:id", (&controller.Role{}).Update)
		api.GET("/permission/menu/tree", (&controller.Permission{}).MenuTree)
		api.GET("/permission/tree", (&controller.Permission{}).Tree)
	}

}
