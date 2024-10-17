package route

import (
	"gex-api/app/controller/admin"
	"gex-api/app/middleware"
	"gex-api/pkg/builder"
)

// 后台路由组
func AdminApi(service *builder.Engine) {

	// 未授权路由组
	api := service.Group("admin/api")
	{
		api.Post("/user/login", (&admin.User{}).Login) // 后台用户登录
	}

	authApi := service.Group("admin/api", middleware.AdminAuthMiddleware)
	{
		// 用户相关
		authApi.Post("/user/create", (&admin.User{}).Create)               // 创建用户
		authApi.Post("/user/update", (&admin.User{}).Update)               // 更新用户
		authApi.Get("/user/delete", (&admin.User{}).Delete)                // 删除用户
		authApi.Get("/user/list", (&admin.User{}).List)                    // 用户列表
		authApi.Get("/user/detail", (&admin.User{}).Detail)                // 用户详情
		authApi.Post("/user/resetPassword", (&admin.User{}).ResetPassword) // 后台用户重置密码
		authApi.Post("/user/bindRole", (&admin.User{}).BindRole)           // 后台用户绑定角色
		authApi.Post("/userroles", (&admin.User{}).Roles)                  // 后台用户所绑定的角色
		authApi.Post("/user/menuTree", (&admin.User{}).MenuTree)           // 登录的用户所拥有的菜单列表树

		// 角色相关
		authApi.Post("/role/create", (&admin.Role{}).Create)                // 创建角色
		authApi.Post("/role/update", (&admin.Role{}).Update)                // 更新角色
		authApi.Get("/role/delete", (&admin.Role{}).Delete)                 // 删除角色
		authApi.Get("/role/list", (&admin.Role{}).List)                     // 角色列表
		authApi.Get("/role/detail", (&admin.Role{}).Detail)                 // 角色详情
		authApi.Get("/role/option", (&admin.Role{}).Option)                 // 角色选项
		authApi.Get("/role/bindMenu", (&admin.Role{}).BindMenu)             // 角色绑定菜单
		authApi.Get("/role/menus", (&admin.Role{}).Menus)                   // 角色所绑定的菜单
		authApi.Get("/role/bindPermission", (&admin.Role{}).BindPermission) // 角色绑定权限
		authApi.Get("/role/permissions", (&admin.Role{}).Permissions)       // 角色所绑定的权限

		// 菜单相关
		authApi.Post("/menu/create", (&admin.Menu{}).Create) // 创建菜单
		authApi.Post("/menu/update", (&admin.Menu{}).Update) // 更新菜单
		authApi.Get("/menu/delete", (&admin.Menu{}).Delete)  // 删除菜单
		authApi.Get("/menu/tree", (&admin.Menu{}).Tree)      // 菜单列表树
		authApi.Get("/menu/detail", (&admin.Menu{}).Detail)  // 菜单详情

		// 权限相关
		authApi.Post("/permission/create", (&admin.Permission{}).Create) // 创建权限
		authApi.Post("/permission/update", (&admin.Permission{}).Update) // 更新权限
		authApi.Get("/permission/delete", (&admin.Permission{}).Delete)  // 删除权限
		authApi.Get("/permission/list", (&admin.Permission{}).List)      // 权限列表
		authApi.Get("/permission/detail", (&admin.Permission{}).Detail)  // 权限详情
	}

}
