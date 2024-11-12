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

		ctx.Set("userId", userClaims.Id)
		ctx.Set("username", userClaims.Username)

		ctx.Next()
	}
}
