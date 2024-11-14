package router

import (
	"isme-go/app/controller"
	"isme-go/app/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRegister(server *gin.Engine) {

	api := server.Group("/api")
	{
		api.POST("/auth/login", (&controller.Auth{}).Login)    // 登录
		api.GET("/auth/captcha", (&controller.Auth{}).Captcha) // 验证码
	}

	api = server.Group("/api", middleware.Authorization())
	{
		// 认证
		api.POST("/auth/current-role/switch/:roleCode", (&controller.Auth{}).SwitchCurrentRole) // 切换角色
		api.POST("/auth/logout", (&controller.Auth{}).Logout)                                   // 退出登录
		api.POST("/auth/password", (&controller.Auth{}).Password)                               // 重置密码

		// 用户
		api.GET("/user/detail", (&controller.User{}).Detail)                      // 用户详情
		api.GET("/user", (&controller.User{}).Page)                               // 用户列表-分页
		api.DELETE("/user/:id", (&controller.User{}).Delete)                      // 删除用户
		api.PATCH("/user/profile/:id", (&controller.User{}).ProfileUpdate)        // 修改用户信息
		api.PATCH("/user/password/reset/:id", (&controller.User{}).PasswordReset) // 重置用户密码
		api.POST("/user", (&controller.User{}).Add)                               // 新增用户
		api.PATCH("/user/:id", (&controller.User{}).Update)                       // 修改用户

		// 权限
		api.GET("/role/permissions/tree", (&controller.Role{}).PermissionsTree)    // 角色列表树-by token
		api.GET("/permission/menu/tree", (&controller.Permission{}).MenuTree)      // 权限树-菜单
		api.GET("/permission/tree", (&controller.Permission{}).Tree)               // 权限树-all
		api.DELETE("/permission/:id", (&controller.Permission{}).Delete)           // 删除权限
		api.POST("/permission", (&controller.Permission{}).Add)                    // 新增权限
		api.PATCH("/permission/:id", (&controller.Permission{}).Update)            // 修改权限
		api.GET("/permission/button/:parentId", (&controller.Permission{}).Button) // 权限按钮

		// 角色
		api.GET("/role/page", (&controller.Role{}).Page)                      // 角色列表-分页
		api.GET("/role", (&controller.Role{}).List)                           // 角色列表-all
		api.PATCH("/role/:id", (&controller.Role{}).Update)                   // 修改角色
		api.POST("/role", (&controller.Role{}).Add)                           // 新增角色
		api.DELETE("/role/:id", (&controller.Role{}).Delete)                  // 删除角色
		api.PATCH("/role/users/add/:id", (&controller.Role{}).UsersAdd)       // 分配角色-批量
		api.PATCH("/role/users/remove/:id", (&controller.Role{}).UsersRemove) // 取消分配角色-批量

	}
}
