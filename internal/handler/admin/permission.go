package admin

import (
	"breeze-api/internal/model"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Permission struct{}

// 添加
func (t *Permission) CreatePermission(ctx *fiber.Ctx) error {

	type request struct {
		Name      string `json:"name"`
		GroupName string `json:"groupName"`
		Path      string `json:"path"`
		Method    string `json:"method"`
		Status    int    `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Path == "" || req.Method == "" {
		return response.Error(ctx, "参数错误")
	}

	permission := (&service.Permission{}).GetPermissionByPathWithMethod(req.Path, req.Method)

	if permission.Id > 0 {
		return response.Error(ctx, "权限已存在")
	}

	err := (&service.Permission{}).CreatePermission(&model.Permission{
		Name:      req.Name,
		GroupName: req.GroupName,
		Path:      req.Path,
		Method:    req.Method,
		Status:    req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新
func (t *Permission) UpdatePermission(ctx *fiber.Ctx) error {

	type request struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		GroupName string `json:"groupName"`
		Path      string `json:"path"`
		Method    string `json:"method"`
		Status    int    `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Permission{}).UpdatePermission(&model.Permission{
		Id:        req.Id,
		Name:      req.Name,
		GroupName: req.GroupName,
		Path:      req.Path,
		Method:    req.Method,
		Status:    req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除
func (t *Permission) DeletePermission(ctx *fiber.Ctx) error {

	type request struct {
		Id int `json:"id"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Permission{}).DeletePermission(req.Id)

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 列表
func (t *Permission) GetPermissionPage(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	size, _ := strconv.Atoi(ctx.Query("size", "10"))

	list, total := (&service.Permission{}).GetPermissionListByPage(page, size)

	return response.Success(ctx, "成功", map[string]interface{}{
		"list":  list,
		"total": total,
	})
}

// 列表
func (t *Permission) GetPermissions(ctx *fiber.Ctx) error {

	list := (&service.Permission{}).GetPermissionList()

	return response.Success(ctx, "成功", list)
}

// 详情
func (t *Permission) GetPermission(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Query("id"))

	permission := (&service.Permission{}).GetPermissionById(id)

	return response.Success(ctx, "成功", permission)
}
