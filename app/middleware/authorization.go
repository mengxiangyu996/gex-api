package middleware

import (
	"isme-go/app/service"
	"isme-go/app/token"
	"isme-go/framework/message"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userClaims, err := token.ParseToken(ctx)
		if err != nil {
			message.Error(ctx, 401, err.Error())
			ctx.Abort()
			return
		}

		permission := (&service.Permission{}).GetDetailByPathAndMethod(ctx.Request.URL.Path, "")
		if permission.Id <= 0 {
			message.Error(ctx, 403, "没有权限")
			ctx.Abort()
			return
		}

		role := (&service.Role{}).GetDetailByCode(userClaims.CurrentRoleCode)
		if role.Id <= 0 {
			message.Error(ctx, 403, "没有权限")
			ctx.Abort()
			return
		}

		if !(&service.RolePermissionsPermission{}).CheckHasPermission(role.Id, permission.Id) {
			message.Error(ctx, 403, "没有权限")
			ctx.Abort()
			return
		}

		ctx.Set("userId", userClaims.UserId)
		ctx.Set("username", userClaims.Username)
		ctx.Set("roleCode", userClaims.CurrentRoleCode)

		ctx.Next()
	}
}
