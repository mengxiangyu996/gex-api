package handler

import (
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type Index struct{}

func (t *Index) Index(ctx *fiber.Ctx) error {
	return response.Success(ctx, "ok", nil)
}