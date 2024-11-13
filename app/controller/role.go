package controller

import (
	"isme-go/app/request"
	"isme-go/app/response"
	"isme-go/app/service"
	"isme-go/framework/message"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Role struct{}

// 角色资源树
func (*Role) PermissionsTree(ctx *gin.Context) {

	userId := ctx.GetInt("userId")

	roleIds := (&service.UserRolesRole{}).GetRoleIdsByUserId(userId)

	permissionIds := (&service.RolePermissionsPermission{}).GetPermissionIdsByRoleIds(roleIds)

	permissions := (&service.Permission{}).GetListByIds(permissionIds, true)

	tree := (&service.Permission{}).ListToTree(permissions, 0)

	message.Success(ctx, map[string]interface{}{
		"data": tree,
	})
}

// 角色分页
func (*Role) Page(ctx *gin.Context) {

	var param request.RolePage

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	rolePages := make([]response.RolePage, 0)

	roles, total := (&service.Role{}).Page(param)

	if len(roles) > 0 {
		for _, role := range roles {
			var rolePage response.RolePage
			rolePage.Role = role
			permissionIds := (&service.RolePermissionsPermission{}).GetPermissionIdsByRoleIds([]int{role.Id})
			rolePage.PermissionIds = permissionIds
			rolePages = append(rolePages, rolePage)
		}
	}

	message.Success(ctx, map[string]interface{}{
		"data": map[string]interface{}{
			"pageData": rolePages,
			"total":    total,
		},
	})
}

// 角色列表
func (*Role) List(ctx *gin.Context) {

	roles := (&service.Role{}).List()

	message.Success(ctx, map[string]interface{}{
		"data": roles,
	})
}

// 添加角色
func (*Role) Add(ctx *gin.Context) {

	var param request.RoleAdd

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	role := (&service.Role{}).GetDetailByCode(param.Code)
	if role.Id > 0 {
		message.Error(ctx, "角色编码已存在")
		return
	}

	if err := (&service.Role{}).Insert(param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}

// 修改角色
func (*Role) Update(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	var param request.RoleUpdate

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	param.Id = id

	if err := (&service.Role{}).Update(param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}
