package response

import (
	"github.com/gin-gonic/gin"
)

// 响应
type Response struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"-"`
}

// 初始化成功响应
func NewSuccess() *Response {

	return &Response{
		Code: 0,
		Msg:  "OK",
		Data: make(map[string]interface{}),
	}
}

// 初始化失败响应
func NewError() *Response {

	return &Response{
		Code: 10500,
		Msg:  "FAILED",
		Data: make(map[string]interface{}),
	}
}

// 设置响应码
func (r *Response) SetCode(code int) *Response {

	r.Code = code

	return r
}

// 设置响应信息
func (r *Response) SetMsg(msg string) *Response {

	r.Msg = msg

	return r
}

// 设置响应数据
func (r *Response) SetData(key string, value interface{}) *Response {

	if key == "code" || key == "msg" {
		return r
	}

	r.Data[key] = value

	return r
}

// 设置分页响应数据
func (r *Response) SetPageData(rows interface{}, total int) *Response {

	r.Data["pageData"] = rows
	r.Data["total"] = total

	return r
}

// 设置响应数据
func (r *Response) SetDataMap(data map[string]interface{}) *Response {

	for key, value := range data {
		if key == "code" || key == "msg" {
			continue
		}

		r.Data[key] = value
	}

	return r
}

// 序列化返回
func (r *Response) Json(ctx *gin.Context) {

	response := gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	}

	for key, value := range r.Data {
		response[key] = value
	}

	ctx.JSON(200, response)
}
