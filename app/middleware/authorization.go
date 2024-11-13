package middleware

import (
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

		ctx.Set("userId", userClaims.UserId)
		ctx.Set("username", userClaims.Username)
		ctx.Set("roleCode", userClaims.CurrentRoleCode)
		
		// todo 权限处理

		ctx.Next()
	}
}
