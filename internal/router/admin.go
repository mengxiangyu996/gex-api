package router

import (
	"breeze-api/internal/handler/admin"
	"breeze-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// 注册 admin 路由
func AadminRegister(app *fiber.App) {

	// 未授权路由组
	api := app.Group("admin/api")
	{
		api.Post("user/login", (&admin.User{}).Login)
	}

	// 授权路由组
	auth := app.Group("admin/api", (&middleware.Api{}).Handle)
	{
		auth.Post("user/createUser", (&admin.User{}).CreateUser)
		auth.Post("user/updateUser", (&admin.User{}).UpdateUser)
		auth.Post("user/deleteUser", (&admin.User{}).DeleteUser)
		auth.Get("user/getUserPage", (&admin.User{}).GetUserPage)
		auth.Get("user/getUser", (&admin.User{}).GetUser)
		auth.Post("user/updatePassword", (&admin.User{}).UpdatePassword)

		auth.Post("role/createRole", (&admin.Role{}).CreateRole)
		auth.Post("role/updateRole", (&admin.Role{}).UpdateRole)
		auth.Post("role/deleteRole", (&admin.Role{}).DeleteRole)
		auth.Get("role/getRolePage", (&admin.Role{}).GetRolePage)
		auth.Get("role/getRole", (&admin.Role{}).GetRole)

		auth.Post("menu/createMenu", (&admin.Menu{}).CreateMenu)
		auth.Post("menu/updateMenu", (&admin.Menu{}).UpdateMenu)
		auth.Post("menu/deleteMenu", (&admin.Menu{}).DeleteMenu)
		auth.Get("menu/getMenuPage", (&admin.Menu{}).GetMenuPage)
		auth.Get("menu/getMenuTree", (&admin.Menu{}).GetMenuTree)
		auth.Get("menu/getMenu", (&admin.Menu{}).GetMenu)

		auth.Post("permission/createPermission", (&admin.Permission{}).CreatePermission)
		auth.Post("permission/updatePermission", (&admin.Permission{}).UpdatePermission)
		auth.Post("permission/deletePermission", (&admin.Permission{}).DeletePermission)
		auth.Get("permission/getPermissionPage", (&admin.Permission{}).GetPermissionPage)
		auth.Get("permission/getPermissions", (&admin.Permission{}).GetPermissions)
		auth.Get("permission/getPermission", (&admin.Permission{}).GetPermission)
	}

	return
}
