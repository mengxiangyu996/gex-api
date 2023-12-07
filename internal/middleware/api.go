package middleware

import (
	"breeze-api/helper/jwt"
	"breeze-api/pkg/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

// 中间件
type Api struct{}

func (t *Api) Handle(ctx *fiber.Ctx) error {

	token := ctx.Get("Token")

	if token == "" {
		return response.Base(ctx, 1002, "请重新登陆", nil)
	}

	user := jwt.Parse(token)
	if user.Id <= 0 || user.Expire.Before(time.Now()) {
		return response.Base(ctx, 1002, "请重新登陆", nil)
	}

	return ctx.Next()
}