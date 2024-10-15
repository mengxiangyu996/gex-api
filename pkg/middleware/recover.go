package middleware

import (
	"fmt"
	"gex-api/pkg/builder"
	"net/http"
	"runtime"
)

// RecoverMiddleware 定义了一个处理恐慌的中间件
func RecoverMiddleware(next builder.HandlerFunc) builder.HandlerFunc {
	return func(ctx *builder.Context) error {
		defer func() {
			if r := recover(); r != nil {
				// 处理恐慌，生成堆栈跟踪
				stack := make([]byte, 4096)
				length := runtime.Stack(stack, true)

				// 打印堆栈跟踪到标准输出
				fmt.Printf("[PANIC RECOVER] %v\nStack Trace:\n%s\n", r, stack[:length])

				// 设置 HTTP 500 响应
				ctx.Response.WriteHeader(http.StatusInternalServerError)
				ctx.Response.Write([]byte("Internal Server Error"))
			}
		}()

		// 调用下一个处理函数，如果出错则返回错误
		return next(ctx)
	}
}
