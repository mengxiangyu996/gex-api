package middleware

import (
	"bytes"
	logfactory "ruoyi-go/internal/log-factory"
	responsewriter "ruoyi-go/internal/response-writer"
	"ruoyi-go/request"
	iputils "ruoyi-go/utils/ip-utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 操作日志
func OperLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		startTime := time.Now()

		ipAddress := iputils.GetAddress(ctx.ClientIP())

		// 创建操作日志工厂实例
		operateLogFactory := logfactory.OperateLogFactory{}

		// 初始化操作日志
		operateLog := operateLogFactory.NewLog().(*logfactory.OperateLog)

		operateLog.SysOperLog = &request.SysOperLogInsertRequest{
			Title:         "测试",
			BusinessType:  0,
			Method:        ctx.Request.Method,
			RequestMethod: ctx.Request.Method,
			OperatorType:  0,
			OperName:      "测试",
			DeptName:      "测试",
			OperUrl:       ctx.Request.URL.Path,
			OperIp:        ipAddress.Ip,
			OperLocation:  ipAddress.Addr,
			OperParam:     "",
			JsonResult:    "",
			Status:        "0",
			ErrorMsg:      "",
			CostTime:      0,
		}

		writer := &responsewriter.ResponseWriter{
			ResponseWriter: ctx.Writer,
			Body:           &bytes.Buffer{},
		}

		ctx.Writer = writer

		ctx.Next()

		operateLog.SysOperLog.CostTime = int(time.Since(startTime).Milliseconds())

		// 记录操作日志
		operateLog.Record()
	}
}
