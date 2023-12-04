package handler

import (
	"breeze-api/internal/service"
	"breeze-api/internal/service/query"
	"breeze-api/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// 用户
type User struct{}

// 详情
func (t *User) Detail(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Query("id"))

	user := (&service.User{}).Detail(query.WithId(id))

	return response.Success(ctx, "ok", user)
}
