package builder

import (
	"errors"
	"gex-api/pkg/tag"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"

	"github.com/labstack/echo"
)

// Context 扩展了echo.Context，添加了自定义功能和字段
type Context struct {
	echoContext echo.Context           // echo上下文
	Engine      *Engine                // 引擎指针
	Request     *http.Request          // 请求
	Response    http.ResponseWriter    // 响应
	Params      map[string]string      // URL Param "user/:id"
	Querys      map[string]interface{} // URL Quary "?key=value"
}

// IsTLS 检查请求是否使用TLS
func (t *Context) IsTLS() bool {
	return t.echoContext.IsTLS()
}

// IsWebSocket 检查请求是否为WebSocket
func (t *Context) IsWebSocket() bool {
	return t.echoContext.IsWebSocket()
}

// Scheme 返回请求的协议（http或https）
func (t *Context) Scheme() string {
	return t.echoContext.Scheme()
}

// RealIP 返回请求的真实IP地址
func (t *Context) RealIP() string {
	return t.echoContext.RealIP()
}

// Path 返回请求的路径
func (t *Context) Path() string {
	return t.echoContext.Path()
}

// SetPath 设置请求的路径
func (t *Context) SetPath(path string) {
	t.echoContext.SetPath(path)
}

// Param 返回指定名称的URL参数
func (t *Context) Param(name string) string {
	return t.echoContext.Param(name)
}

// ParamNames 返回所有URL参数的名称
func (t *Context) ParamNames() []string {
	return t.echoContext.ParamNames()
}

// SetParamNames 设置URL参数的名称
func (t *Context) SetParamNames(names ...string) {
	t.echoContext.SetParamNames(names...)
}

// ParamValues 返回所有URL参数的值
func (t *Context) ParamValues() []string {
	return t.echoContext.ParamValues()
}

// SetParamValues 设置URL参数的值
func (t *Context) SetParamValues(value ...string) {
	t.echoContext.SetParamValues(value...)
}

// QueryParam 返回指定名称的查询参数
func (t *Context) QueryParam(name string) string {
	return t.echoContext.QueryParam(name)
}

// QueryParams 返回所有查询参数
func (t *Context) QueryParams() url.Values {
	return t.echoContext.QueryParams()
}

// QueryString 返回原始查询字符串
func (t *Context) QueryString() string {
	return t.echoContext.QueryString()
}

// FormValue 返回指定名称的表单值
func (t *Context) FormValue(name string) string {
	return t.echoContext.FormValue(name)
}

// FormParams 返回所有表单参数
func (t *Context) FormParams() (url.Values, error) {
	return t.echoContext.FormParams()
}

// FormFile 返回指定名称的文件表单字段
func (t *Context) FormFile(name string) (*multipart.FileHeader, error) {
	return t.echoContext.FormFile(name)
}

// MultipartForm 返回multipart表单
func (t *Context) MultipartForm() (*multipart.Form, error) {
	return t.echoContext.MultipartForm()
}

// Cookie 返回指定名称的cookie
func (t *Context) Cookie(name string) (*http.Cookie, error) {
	return t.echoContext.Cookie(name)
}

// SetCookie 设置cookie
func (t *Context) SetCookie(cookie *http.Cookie) {
	t.echoContext.SetCookie(cookie)
}

// Cookies 返回所有cookies
func (t *Context) Cookies() []*http.Cookie {
	return t.echoContext.Cookies()
}

// Get 返回指定键的值
func (t *Context) Get(key string) interface{} {
	return t.echoContext.Get(key)
}

// Set 设置指定键的值
func (t *Context) Set(key string, value interface{}) {
	t.echoContext.Set(key, value)
}

// Bind 绑定请求数据到指定结构体
func (t *Context) Bind(i interface{}) error {
	return t.echoContext.Bind(i)
}

// Validate 验证请求数据
func (t *Context) Validate(i interface{}) error {
	return t.echoContext.Validate(i)
}

// Render 渲染指定模板并返回HTML响应
func (t *Context) Render(code int, name string, data interface{}) error {
	return t.echoContext.Render(code, name, data)
}

// HTML 返回HTML响应
func (t *Context) HTML(code int, html string) error {
	return t.echoContext.HTML(code, html)
}

// HTMLBlob 返回HTML响应，使用字节数组
func (t *Context) HTMLBlob(code int, b []byte) error {
	return t.echoContext.HTMLBlob(code, b)
}

// String 返回文本响应
func (t *Context) String(code int, s string) error {
	return t.echoContext.String(code, s)
}

// JSON 返回JSON响应
func (t *Context) JSON(code int, i interface{}) error {
	return t.echoContext.JSON(code, i)
}

// JSONPretty 返回格式化的JSON响应
func (t *Context) JSONPretty(code int, i interface{}, indent string) error {
	return t.echoContext.JSONPretty(code, i, indent)
}

// JSONBlob 返回JSON响应，使用字节数组
func (t *Context) JSONBlob(code int, b []byte) error {
	return t.echoContext.JSONBlob(code, b)
}

// JSONP 返回JSONP响应
func (t *Context) JSONP(code int, callback string, i interface{}) error {
	return t.echoContext.JSONP(code, callback, i)
}

// JSONPBlob 返回JSONP响应，使用字节数组
func (t *Context) JSONPBlob(code int, callback string, b []byte) error {
	return t.echoContext.JSONPBlob(code, callback, b)
}

// XML 返回XML响应
func (t *Context) XML(code int, i interface{}) error {
	return t.echoContext.XML(code, i)
}

// XMLPretty 返回格式化的XML响应
func (t *Context) XMLPretty(code int, i interface{}, indent string) error {
	return t.echoContext.XMLPretty(code, i, indent)
}

// XMLBlob 返回XML响应，使用字节数组
func (t *Context) XMLBlob(code int, b []byte) error {
	return t.echoContext.XMLBlob(code, b)
}

// Blob 返回二进制数据响应
func (t *Context) Blob(code int, contentType string, b []byte) error {
	return t.echoContext.Blob(code, contentType, b)
}

// Stream 返回流式响应
func (t *Context) Stream(code int, contentType string, r io.Reader) error {
	return t.echoContext.Stream(code, contentType, r)
}

// Attachment 返回附件响应
func (t *Context) Attachment(file, name string) error {
	return t.echoContext.Attachment(file, name)
}

// Inline 返回内联响应
func (t *Context) Inline(file, name string) error {
	return t.echoContext.Inline(file, name)
}

// NoContent 返回无内容响应
func (t *Context) NoContent(code int) error {
	return t.echoContext.NoContent(code)
}

// Redirect 返回重定向响应
func (t *Context) Redirect(code int, url string) error {
	return t.echoContext.Redirect(code, url)
}

// Error 记录错误并返回错误响应
func (t *Context) Error(err error) {
	t.echoContext.Error(err)
}

// Reset 重置上下文，通常在中间件中使用
func (t *Context) Reset(request *http.Request, response http.ResponseWriter) {
	t.echoContext.Reset(request, response)
}

// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓ 扩展方法 ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓

// GetEchoContext 返回Echo框架上下文
func (t *Context) GetEchoContext() echo.Context {
	return t.echoContext
}

// Method 返回请求方式
func (t *Context) Method() string {
	return t.echoContext.Request().Method
}

// GetHeader 返回Header值
func (t *Context) GetHeader(key string) string {
	return t.echoContext.Request().Header.Get(key)
}

// SetHeader 设置Header值
func (t *Context) SetHeader(key string, value string) {
	t.echoContext.Request().Header.Set(key, value)
}

// BindX Bind的增强
func (t *Context) BindX(i interface{}) error {

	if err := t.echoContext.Bind(i); err != nil {
		return err
	}

	// Set tag defaults
	// type Example struct {
	// 	Page      int                    `default:"1"`
	// 	Size      int                    `default:"10"`
	// 	IsBool    bool                   `default:"true"`
	// 	Keyword   string                 `default:"hello world"`
	// 	Number    float64                `default:"1.23"`
	// 	Status    [3]int                 `default:"1,2"`
	// 	TimeRange []string               `default:"12:30,13:30"`
	// 	Data      map[string]interface{} `default:"age:18,name:zero"`
	// }
	// {Page:1 Size:10 Keyword:hello world Number:1.23 Status:[1 2 0] TimeRange:[12:30 13:30] Data:map[age:18 name:zero]}
	tag.SetDefaults(reflect.ValueOf(i).Elem())

	return nil
}

// Success 发送成功响应，带有自定义消息和可选数据
//
// Success("ok")
//
// Success(map[string]interface{}{ "date": "2006-01-02 15:04:05" })
//
// Success("ok", map[string]interface{}{ "date": "2006-01-02 15:04:05" })
func (t *Context) Success(arg ...interface{}) error {

	var message string
	var data interface{}
	var ok bool

	switch len(arg) {
	case 1:
		message, ok = arg[0].(string)
		if !ok {
			message = "success"
			data = arg[0]
		}
	case 2:
		message, ok = arg[0].(string)
		if !ok {
			return errors.New("first argument must be a string")
		}
		data = arg[1]
	default:
		message = "success"
	}

	return t.JSON(200, json(10200, message, data))
}

// Fail 发送失败响应，带有自定义消息和可选数据
//
// Fail("fail")
//
// Fail(map[string]interface{}{ "date": "2006-01-02 15:04:05" })
//
// Fail("fail", map[string]interface{}{ "date": "2006-01-02 15:04:05" })
func (t *Context) Fail(arg ...interface{}) error {

	var message string
	var data interface{}
	var ok bool

	switch len(arg) {
	case 1:
		message, ok = arg[0].(string)
		if !ok {
			message = "fail"
			data = arg[0]
		}
	case 2:
		message, ok = arg[0].(string)
		if !ok {
			return errors.New("first argument must be a string")
		}
		data = arg[1]
	default:
		message = "fail"
	}

	return t.JSON(200, json(10500, message, data))
}

// Json 发送json响应，带有自定义消息和可选数据
func (t *Context) Json(code int, message string, data ...interface{}) error {
	return t.JSON(200, json(code, message, data))
}
