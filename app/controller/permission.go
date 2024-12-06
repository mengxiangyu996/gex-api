package controller

import (
	"isme-go/app/dto"
	"isme-go/app/service"
	"isme-go/framework/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 资源
type Permission struct{}

// 获取资源树
func (*Permission) MenuTree(ctx *gin.Context) {

	list := (&service.Permission{}).GetList(false)

	tree := (&service.Permission{}).ListToTree(list, 0)

	response.NewSuccess().SetData("data", tree).Json(ctx)
}

// 获取资源树
func (*Permission) Tree(ctx *gin.Context) {

	list := (&service.Permission{}).GetList(true)

	tree := (&service.Permission{}).ListToTree(list, 0)

	response.NewSuccess().SetData("data", tree).Json(ctx)
}

// 添加资源
func (*Permission) Add(ctx *gin.Context) {

	var param dto.PermissionAddRequest

	if err := ctx.Bind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if param.Code == "" {
		response.NewError().SetMsg("编码不能为空").Json(ctx)
		return
	}

	if err := (&service.Permission{}).Add(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除资源
func (*Permission) Delete(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := (&service.Permission{}).Delete(id); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改资源
func (*Permission) Update(ctx *gin.Context) {

	var param dto.PermissionUpdateRequest

	if err := ctx.Bind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	param.Id, _ = strconv.Atoi(ctx.Param("id"))

	if err := (&service.Permission{}).Update(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 获取权限按钮
func (*Permission) Button(ctx *gin.Context) {

	parentId, _ := strconv.Atoi(ctx.Param("parentId"))

	buttons := (&service.Permission{}).GetButtons(parentId)

	response.NewSuccess().SetData("data", buttons).Json(ctx)
}
