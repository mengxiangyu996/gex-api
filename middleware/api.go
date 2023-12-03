package middleware

import "github.com/gofiber/fiber/v2"

// 中间件
type Api struct{}

func (t *Api) Handle(ctx *fiber.Ctx) error {

	// TODO

	return ctx.Next()
}