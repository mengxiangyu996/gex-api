package response

import "github.com/gofiber/fiber/v2"

type Response struct{}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 成功
func Success(ctx *fiber.Ctx, message string, data interface{}) error {
	return Base(ctx, 0, message, data)
}

// 失败
func Error(ctx *fiber.Ctx, message string) error {
	return Base(ctx, 1001, message, nil)
}

// 自定义
func Base(ctx *fiber.Ctx, code int, message string, data interface{}) error {
	return ctx.JSON(result(code, message, data))
}

func result(code int, message string, data interface{}) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
