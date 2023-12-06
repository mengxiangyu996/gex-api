package admin

import (
	"breeze-api/helper/encrypt"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type User struct{}

// 登录
func (t *User) Login(ctx *fiber.Ctx) error {

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Username == "" || req.Password == "" {
		return response.Error(ctx, "参数错误")
	}

	user := (&service.User{}).GetUserByUsername(req.Username)

	if user.Status != 1 {
		return response.Error(ctx, "账号异常")
	}

	if !encrypt.Compare(user.Password, req.Password) {
		return response.Error(ctx, "密码错误")
	}

	return response.Success(ctx, "ok", nil)
}