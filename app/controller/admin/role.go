package admin

import (
	"gex-api/app/internal/utils"
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

	param.IsPaging = true
	
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

// 角色选项
func (*Role) Option(ctx *builder.Context) error {

	option := (&service.Role{}).Option()

	return ctx.Success(option)
}

// 角色绑定菜单
func (*Role) BindMenu(ctx *builder.Context) error {

	var param request.RoleBindMenu

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if err := (&service.Role{}).BindMenu(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 角色所绑定的菜单
func (*Role) Menus(ctx *builder.Context) error {

	var param request.QueryId

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	menus := (&service.Menu{}).GetListByIds(nil)
	menuIds := (&service.Role{}).GetBindMenu(param.Id)

	for _, menu := range menus {
		if utils.Contains(menuIds, menu.Id) {
			menu.IsBind = true
		}
	}

	tree := (&service.Menu{}).ListToTree(menus, 0)

	return ctx.Success(tree)
}

// 角色绑定权限
func (*Role) BindPermission(ctx *builder.Context) error {

	var param request.RoleBindPermission

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	if err := (&service.Role{}).BindPermission(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	return ctx.Success()
}

// 角色所绑定的权限
func (*Role) Permissions(ctx *builder.Context) error {

	var param request.RoleBindPermission

	if err := ctx.BindX(&param); err != nil {
		return ctx.Fail(err.Error())
	}

	permissions := (&service.Permission{}).GetListByIds(nil)

	permissionIds := (&service.Role{}).GetBindPermission(param.RoleId)

	for _, permission := range permissions {
		if utils.Contains(permissionIds, permission.Id) {
			permission.IsBind = true
		}
	}

	return ctx.Success(permissions)
}
