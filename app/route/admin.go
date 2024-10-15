package route

import (
	"gex-api/app/controller"
	"gex-api/app/middleware"
	"gex-api/pkg/builder"
)

// 后台路由组
func AdminApi(service *builder.Engine) {

	// 未授权路由组
	api := service.Group("admin/api")
	{
		api.Post("user/login", (&controller.User{}).Login) // 后台用户登录
	}

	authApi := service.Group("admin/api", middleware.AdminAuthMiddleware)
	{
		// 用户相关
		authApi.Post("user/create", (&controller.User{}).Create)               // 创建用户
		authApi.Post("user/update", (&controller.User{}).Update)               // 更新用户
		authApi.Get("user/delete", (&controller.User{}).Delete)                // 删除用户
		authApi.Get("user/list", (&controller.User{}).List)                    // 用户列表
		authApi.Get("user/detail", (&controller.User{}).Detail)                // 用户详情
		authApi.Post("user/resetPassword", (&controller.User{}).ResetPassword) // 后台用户重置密码
		authApi.Post("user/bindRole", (&controller.User{}).BindRole)           // 后台用户绑定角色

		// 角色相关
		authApi.Post("role/create", (&controller.Role{}).Create) // 创建角色
		authApi.Post("role/update", (&controller.Role{}).Update) // 更新角色
		authApi.Get("role/delete", (&controller.Role{}).Delete)  // 删除角色
		authApi.Get("role/list", (&controller.Role{}).List)      // 角色列表
		authApi.Get("role/detail", (&controller.Role{}).Detail)  // 角色详情

		// 菜单相关
		authApi.Post("menu/create", (&controller.Menu{}).Create) // 创建菜单
		authApi.Post("menu/update", (&controller.Menu{}).Update) // 更新菜单
		authApi.Get("menu/delete", (&controller.Menu{}).Delete)  // 删除菜单
		authApi.Get("menu/tree", (&controller.Menu{}).Tree)      // 菜单列表树
		authApi.Get("menu/detail", (&controller.Menu{}).Detail)  // 菜单详情

		// 权限相关
		authApi.Post("permission/create", (&controller.Permission{}).Create) // 创建权限
		authApi.Post("permission/update", (&controller.Permission{}).Update) // 更新权限
		authApi.Get("permission/delete", (&controller.Permission{}).Delete)  // 删除权限
		authApi.Get("permission/list", (&controller.Permission{}).List)      // 权限列表
		authApi.Get("permission/detail", (&controller.Permission{}).Detail)  // 权限详情
	}

}
