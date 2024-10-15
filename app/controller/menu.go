package controller

import (
	"gex-api/app/request"
	"gex-api/app/service"
	"gex-api/pkg/builder"
)

// 菜单
type Menu struct{}

// 创建菜单
func (*Menu) Create(ctx *builder.Context) error {

	var param request.CreateMenu

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if param.Name == "" || param.Type <= 0 || param.Path == "" {
		return ctx.Fail("参数错误")
	}

	if err := (&service.Menu{}).Create(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 更新菜单
func (*Menu) Update(ctx *builder.Context) error {

	var param request.UpdateMenu

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if err := (&service.Menu{}).Update(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 删除菜单
func (*Menu) Delete(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if (&service.Menu{}).IsExistChildren(param.Id) {
		return ctx.Fail("存在下级菜单")
	}

	if err := (&service.Menu{}).DeleteById(param.Id); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 菜单列表树
func (*Menu) Tree(ctx *builder.Context) error {

	list := (&service.Menu{}).GetListByIds(nil)

	tree := (&service.Menu{}).ListToTree(list, 0)

	return ctx.Success(tree)
}

// 菜单详情
func (*Menu) Detail(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	detail := (&service.Menu{}).GetDetailById(param.Id)

	return ctx.Success(detail)
}
