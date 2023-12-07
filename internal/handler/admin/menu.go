package admin

import (
	"breeze-api/internal/model"
	"breeze-api/internal/service"
	"breeze-api/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Menu struct{}

// 添加
func (t *Menu) CreateMenu(ctx *fiber.Ctx) error {

	type request struct {
		ParentId  int    `json:"parentId"`
		Name      string `json:"name"`
		Type      int    `json:"type"`
		Sort      int    `json:"sort"`
		Path      string `json:"path"`
		Component string `json:"component"`
		Icon      string `json:"icon"`
		Redirect  string `json:"redirect"`
		Status    int    `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Name == "" || req.Type <= 0 || req.Path == "" {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Menu{}).CreateMenu(&model.Menu{
		ParentId:  req.ParentId,
		Name:      req.Name,
		Type:      req.Type,
		Sort:      req.Sort,
		Path:      req.Path,
		Component: req.Component,
		Icon:      req.Icon,
		Redirect:  req.Redirect,
		Status:    req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 更新
func (t *Menu) UpdateMenu(ctx *fiber.Ctx) error {

	type request struct {
		Id        int    `json:"id"`
		ParentId  int    `json:"parentId"`
		Name      string `json:"name"`
		Type      int    `json:"type"`
		Sort      int    `json:"sort"`
		Path      string `json:"path"`
		Component string `json:"component"`
		Icon      string `json:"icon"`
		Redirect  string `json:"redirect"`
		Status    int    `json:"status"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Menu{}).UpdateMenu(&model.Menu{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Name:      req.Name,
		Type:      req.Type,
		Sort:      req.Sort,
		Path:      req.Path,
		Component: req.Component,
		Icon:      req.Icon,
		Redirect:  req.Redirect,
		Status:    req.Status,
	})

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 删除
func (t *Menu) DeleteMenu(ctx *fiber.Ctx) error {

	type request struct {
		Id int `json:"id"`
	}

	var req request

	ctx.BodyParser(&req)

	if req.Id <= 0 {
		return response.Error(ctx, "参数错误")
	}

	err := (&service.Menu{}).DeleteMenu(req.Id)

	if err != nil {
		return response.Error(ctx, "失败")
	}

	return response.Success(ctx, "成功", nil)
}

// 列表
func (t *Menu) GetMenuPage(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	size, _ := strconv.Atoi(ctx.Query("size", "10"))

	list, total := (&service.Menu{}).GetMenuListByPage(page, size)

	return response.Success(ctx, "成功", map[string]interface{}{
		"list":  list,
		"total": total,
	})
}

// 菜单列表
func (t *Menu) GetMenuTree(ctx *fiber.Ctx) error {

	list := (&service.Menu{}).GetMenuTree()

	return response.Success(ctx, "成功", list)
}

// 详情
func (t *Menu) GetMenu(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Query("id"))

	menu := (&service.Menu{}).GetMenuById(id)

	return response.Success(ctx, "成功", menu)
}
