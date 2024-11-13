package controller

import (
	"isme-go/app/request"
	"isme-go/app/service"
	"isme-go/framework/message"

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
		message.Error(ctx, "")
	}

	message.Success(ctx)
}
