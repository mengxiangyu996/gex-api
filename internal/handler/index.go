package handler

import (
	"breeze-api/helper/upload"
	"breeze-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type Index struct{}

func (t *Index) Index(ctx *fiber.Ctx) error {

	f, _ := ctx.FormFile("file")

	upload.File(f)

	return response.Success(ctx, "ok", nil)
}