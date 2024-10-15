package controller

import (
	"gex-api/app/request"
	"gex-api/app/response"
	"gex-api/app/service"
	"gex-api/pkg/builder"
)

// 角色
type Role struct{}

// 创建角色
func (*Role) Create(ctx *builder.Context) error {

	var param request.CreateRole

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if err := (&service.Role{}).Create(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 更新角色
func (*Role) Update(ctx *builder.Context) error {

	var param request.UpdateRole

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if err := (&service.Role{}).Update(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 删除角色
func (*Role) Delete(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if err := (&service.Role{}).DeleteById(param.Id); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 角色列表
func (*Role) List(ctx *builder.Context) error {

	var param request.QueryListRole

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	list, count := (&service.Role{}).GetList(&param)

	return ctx.Success(&response.List{
		List:  list,
		Total: count,
	})
}

// 角色详情
func (*Role) Detail(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	detail := (&service.Role{}).GetDetailById(param.Id)

	return ctx.Success(detail)
}
