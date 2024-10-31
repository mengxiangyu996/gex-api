package middleware

import (
	"ruoyi-go/app/token"
	"ruoyi-go/framework/message"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userClaims, err := token.ParseToken(ctx)
		if err != nil {
			message.Error(ctx, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("userId", userClaims.UserId)
		ctx.Set("userName", userClaims.UserName)

		ctx.Next()
	}
}
