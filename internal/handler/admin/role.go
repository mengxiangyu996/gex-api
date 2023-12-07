package admin

import (
	"breeze-api/internal/model"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Role struct{}

// 添加
func (t *Role) CreateRole(ctx *fiber.Ctx) error {

	type request struct {
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Name == "" {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Role{}).CreateRole(&model.Role{
		Name:   req.Name,
		Status: req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新
func (t *Role) UpdateRole(ctx *fiber.Ctx) error {

	type request struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Role{}).UpdateRole(&model.Role{
		Id:     req.Id,
		Name:   req.Name,
		Status: req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除
func (t *Role) DeleteRole(ctx *fiber.Ctx) error {

	type request struct {
		Id int `json:"id"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Role{}).DeleteRole(req.Id)

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 列表
func (t *Role) GetRolePage(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	size, _ := strconv.Atoi(ctx.Query("size", "10"))

	list, total := (&service.Role{}).GetRoleListByPage(page, size)

	return response.Success(ctx, "成功", map[string]interface{}{
		"list":  list,
		"total": total,
	})
}

// 详情
func (t *Role) GetRole(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Query("id"))

	role := (&service.Role{}).GetRoleById(id)

	return response.Success(ctx, "成功", role)
}
