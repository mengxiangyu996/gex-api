package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 成功响应
func Success(c *gin.Context, args ...interface{}) {

	code := SUCCESS
	msg := getMessage(code)

	response := gin.H{
		"code": code,
		"message":  msg,
	}

	for _, arg := range args {
		if v, ok := arg.(int); ok {
			response["code"] = v
			response["message"] = getMessage(v)
		}

		if v, ok := arg.(string); ok {
			response["message"] = v
		}

		if v, ok := arg.(map[string]interface{}); ok {
			for key, value := range v {
				response[key] = value
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

// 失败响应
func Error(c *gin.Context, args ...interface{}) {

	code := ERROR_CODE
	msg := getMessage(code)

	response := gin.H{
		"code": code,
		"message":  msg,
	}

	for _, arg := range args {
		if v, ok := arg.(int); ok {
			response["code"] = v
			response["message"] = getMessage(v)
		}

		if v, ok := arg.(string); ok {
			response["message"] = v
		}

		if v, ok := arg.(map[string]interface{}); ok {
			for key, value := range v {
				response[key] = value
			}
		}
	}

	c.JSON(http.StatusOK, response)
}
