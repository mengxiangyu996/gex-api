package controller

import (
	"isme-go/app/dto"
	"isme-go/app/service"
	"isme-go/framework/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Role struct{}

// 角色资源树
func (*Role) PermissionsTree(ctx *gin.Context) {

	userId := ctx.GetInt("userId")
	roleCode := ctx.GetString("roleCode")

	roleIds := (&service.UserRolesRole{}).GetRoleIdsByUserId(userId)

	permissionIds := (&service.RolePermissionsPermission{}).GetPermissionIdsByRoleIds(roleIds)

	if roleCode == "SUPER_ADMIN" {
		permissionIds = []int{}
	}

	permissions := (&service.Permission{}).GetListByIds(permissionIds, true)

	tree := (&service.Permission{}).ListToTree(permissions, 0)

	response.NewSuccess().SetData("data", tree).Json(ctx)
}

// 角色分页
func (*Role) Page(ctx *gin.Context) {

	var param dto.RolePageRequest

	if err := ctx.Bind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	rolePages := make([]dto.RolePageResponse, 0)

	roles, total := (&service.Role{}).Page(param)

	if len(roles) > 0 {
		for _, role := range roles {
			var rolePage dto.RolePageResponse
			rolePage.RoleResponse = role
			permissionIds := (&service.RolePermissionsPermission{}).GetPermissionIdsByRoleIds([]int{role.Id})
			rolePage.PermissionIds = permissionIds
			rolePages = append(rolePages, rolePage)
		}
	}

	response.NewSuccess().SetData("data", map[string]interface{}{
		"pageData": rolePages,
		"total":    total,
	}).Json(ctx)
}

// 角色列表
func (*Role) List(ctx *gin.Context) {

	roles := (&service.Role{}).List()

	response.NewSuccess().SetData("data", roles).Json(ctx)
}

// 添加角色
func (*Role) Add(ctx *gin.Context) {

	var param dto.RoleAddRequest

	if err := ctx.Bind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	role := (&service.Role{}).GetDetailByCode(param.Code)
	if role.Id > 0 {
		response.NewError().SetMsg("角色编码已存在").Json(ctx)
		return
	}

	if err := (&service.Role{}).Insert(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改角色
func (*Role) Update(ctx *gin.Context) {

	var param dto.RoleUpdateRequest

	if err := ctx.Bind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	param.Id, _ = strconv.Atoi(ctx.Param("id"))

	role := (&service.Role{}).GetDetailById(param.Id)

	if role.Code == "SUPER_ADMIN" {
		response.NewError().SetMsg("超级管理员角色不能修改").Json(ctx)
		return
	}

	if err := (&service.Role{}).Update(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除角色
func (*Role) Delete(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	role := (&service.Role{}).GetDetailById(id)
	if role.Code == "SUPER_ADMIN" {
		response.NewError().SetMsg("超级管理员角色不能删除").Json(ctx)
		return
	}

	if err := (&service.Role{}).Delete(id); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 取消分配角色-批量
func (*Role) UsersRemove(ctx *gin.Context) {

	var param dto.RoleUsersRemoveRequest

	if err := ctx.Bind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	param.RoleId, _ = strconv.Atoi(ctx.Param("id"))

	if err := (&service.UserRolesRole{}).Delete(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 分配角色-批量
func (*Role) UsersAdd(ctx *gin.Context) {

	var param dto.RoleUsersAddRequest

	if err := ctx.Bind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	param.RoleId, _ = strconv.Atoi(ctx.Param("id"))

	if err := (&service.UserRolesRole{}).Insert(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
