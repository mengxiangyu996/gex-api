package middleware

import (
	"gex-api/app/internal/utils"
	"gex-api/app/service"
	"gex-api/pkg/builder"
)

// 后台鉴权
func AdminAuthMiddleware(next builder.HandlerFunc) builder.HandlerFunc {
	return func(ctx *builder.Context) error {

		id, err := utils.ParseTokenPayload(ctx.GetToken())
		if err != nil {
			return ctx.Json(10401, err.Error())
		}

		user := (&service.User{}).GetDetailById(id)
		if user.Id == 1 && user.Role == 2 {
			return next(ctx) // 超级管理员直接通过
		}

		// 非超级管理员鉴权
		roleIds := (&service.User{}).GetBindRole(user.Id)
		if len(roleIds) <= 0 {
			return ctx.Json(10401, "未绑定角色")
		}

		permission := (&service.Permission{}).GetDetailByPathAndMethod(ctx.Path(), ctx.Method())
		if permission.Id <= 0 || permission.Status != 1 {
			return ctx.Json(10401, "权限不存在") // 访问的权限不存在
		}

		// 遍历所有绑定的角色拥有的权限是否包含访问的权限
		for _, roleId := range roleIds {
			permissionIds := (&service.Role{}).GetBindPermission(roleId)
			for _, permissionId := range permissionIds {
				// 权限存在，通过鉴权，允许请求
				if permissionId == permission.Id {
					return next(ctx)
				}
			}
		}

		return ctx.Json(10401, "无权限")
	}
}
