package controller

import (
	"isme-go/app/service"
	"isme-go/framework/message"

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
		"data":      tree,
	})
}
