package controller

import (
	"isme-go/app/request"
	"isme-go/app/service"
	"isme-go/framework/message"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 资源
type Permission struct{}

// 获取资源树
func (*Permission) MenuTree(ctx *gin.Context) {

	list := (&service.Permission{}).GetList(false)

	tree := (&service.Permission{}).ListToTree(list, 0)

	message.Success(ctx, map[string]interface{}{
		"data": tree,
	})
}

// 获取资源树
func (*Permission) Tree(ctx *gin.Context) {

	list := (&service.Permission{}).GetList(true)

	tree := (&service.Permission{}).ListToTree(list, 0)

	message.Success(ctx, map[string]interface{}{
		"data": tree,
	})
}

// 添加资源
func (*Permission) Add(ctx *gin.Context) {

	var param request.PermissionAdd

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	if param.Code == "" {
		message.Error(ctx, "编码不能为空")
		return
	}

	if err := (&service.Permission{}).Add(param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}

// 删除资源
func (*Permission) Delete(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := (&service.Permission{}).Delete(id); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}

// 修改资源
func (*Permission) Update(ctx *gin.Context) {

	var param request.PermissionUpdate

	if err := ctx.Bind(&param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	param.Id, _ = strconv.Atoi(ctx.Param("id"))

	if err := (&service.Permission{}).Update(param); err != nil {
		message.Error(ctx, err.Error())
		return
	}

	message.Success(ctx)
}

// 获取权限按钮
func (*Permission) Button(ctx *gin.Context) {

	parentId, _ := strconv.Atoi(ctx.Param("parentId"))

	buttons := (&service.Permission{}).GetButtons(parentId)

	message.Success(ctx, map[string]interface{}{
		"data": buttons,
	})
}
