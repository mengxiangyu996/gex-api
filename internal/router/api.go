package router

import (
	"breeze-api/internal/handler"
	"breeze-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// 注册Api路由
func ApiRegister(app *fiber.App) {

	api := app.Group("api", (&middleware.Api{}).Handle)
	{
		api.Get("index", (&handler.Index{}).Index)
	}

	return
}