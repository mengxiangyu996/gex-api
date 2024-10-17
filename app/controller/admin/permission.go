package admin

import (
	"gex-api/app/request"
	"gex-api/app/response"
	"gex-api/app/service"
	"gex-api/pkg/builder"
)

// 权限
type Permission struct{}

// 创建权限
func (*Permission) Create(ctx *builder.Context) error {

	var param request.CreatePermission

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if param.Path == "" || param.Method == "" {
		return ctx.Fail("参数错误")
	}

	if permission := (&service.Permission{}).GetDetailByPathAndMethod(param.Path, param.Method); permission.Id > 0 {
		return ctx.Fail("权限已存在")
	}

	if err := (&service.Permission{}).Create(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 更新权限
func (*Permission) Update(ctx *builder.Context) error {

	var param request.UpdatePermission

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if param.Id <= 0 || param.Path == "" || param.Method == "" {
		return ctx.Fail("参数错误")
	}

	if permission := (&service.Permission{}).GetDetailByPathAndMethod(param.Path, param.Method); permission.Id > 0 && param.Id != permission.Id {
		return ctx.Fail("权限已存在")
	}

	if err := (&service.Permission{}).Update(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 删除权限
func (*Permission) Delete(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if err := (&service.Permission{}).DeleteById(param.Id); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 权限列表
func (*Permission) List(ctx *builder.Context) error {

	var param request.QueryListPermission

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	param.IsPaging = true

	list, count := (&service.Permission{}).GetList(&param)

	return ctx.Success(&response.List{
		List:  list,
		Total: count,
	})
}

// 权限详情
func (*Permission) Detail(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	detail := (&service.Permission{}).GetDetailById(param.Id)

	return ctx.Success(detail)
}
