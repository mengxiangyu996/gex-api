package router

import (
	"breeze-api/handler"
	"breeze-api/middleware"

	"github.com/gofiber/fiber/v2"
)

// 注册Api路由
func ApiRegister(app *fiber.App) {

	api := app.Group("api", (&middleware.Api{}).Handle)
	{
		api.Get("user/detail", (&handler.User{}).Detail)
	}

	return
}