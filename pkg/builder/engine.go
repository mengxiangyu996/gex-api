package builder

import (
	"context"
	"gex-api/pkg/dal"
	"net/http"

	"github.com/labstack/echo"
)

// HandlerFunc 定义了一个处理函数类型，接收自定义的Context
type HandlerFunc func(*Context) error

// MiddlewareFunc 定义了一个中间件函数类型，接收一个处理函数并返回一个处理函数
type MiddlewareFunc func(HandlerFunc) HandlerFunc

// Engine 是应用的核心结构，封装了Echo框架的实例和其他配置
type Engine struct {
	echo   *echo.Echo
	config *Config
	*RouteGroup
}

// RouteGroup 表示一组路由，可以应用特定的中间件
type RouteGroup struct {
	engine    *Engine
	echoGroup *echo.Group
}

// Config 存储应用的配置信息
type Config struct {
	GormConfig  *dal.GormConfig
	RedisConfig *dal.RedisConfig
}

// New 初始化并返回一个新的Engine实例
func New(config *Config) *Engine {
	// 初始化echo
	e := echo.New()

	// 隐藏启动横幅
	e.HideBanner = true

	// 初始化数据访问层
	dal.Init(&dal.DBConfig{
		GormConfig:  config.GormConfig,
		RedisConfig: config.RedisConfig,
	})

	engine := &Engine{
		echo:   e,
		config: config,
	}

	engine.RouteGroup = &RouteGroup{
		engine: engine,
	}

	return engine
}

// GetConfig 返回引擎的配置
func (t *Engine) GetConfig() *Config {
	return t.config
}

// GetEcho 返回引擎的Echo实例
func (t *Engine) GetEcho() *echo.Echo {
	return t.echo
}

// NewContext 创建并返回一个新的Context实例
func (t *Engine) NewContext(request *http.Request, response http.ResponseWriter) *Context {
	return &Context{
		echoContext: t.echo.NewContext(request, response),
		Engine:      t,
		Request:     request,
		Response:    response,
	}
}

// Static 注册一个静态文件处理器
func (t *Engine) Static(prefix, root string) {
	t.echo.Static(prefix, root)
}

// Use 添加全局中间件
func (t *Engine) Use(middlewares ...MiddlewareFunc) {
	t.echo.Use(t.handleMiddlewares(middlewares)...)
}

// Group 创建一个新的路由组，并应用中间件
func (t *Engine) Group(prefix string, middlewares ...MiddlewareFunc) *RouteGroup {

	echoGroup := t.echo.Group(prefix, t.handleMiddlewares(middlewares)...)

	routeGroup := &RouteGroup{
		engine:    t,
		echoGroup: echoGroup,
	}

	return routeGroup
}

// Get 注册一个GET请求的路由
func (t *Engine) Get(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echo.GET(path, t.wrapHandler(handler), t.handleMiddlewares(middlewares)...)

	return nil
}

// Post 注册一个POST请求的路由
func (t *Engine) Post(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echo.POST(path, t.wrapHandler(handler), t.handleMiddlewares(middlewares)...)

	return nil
}

// Put 注册一个PUT请求的路由
func (t *Engine) Put(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echo.PUT(path, t.wrapHandler(handler), t.handleMiddlewares(middlewares)...)

	return nil
}

// Delete 注册一个DELETE请求的路由
func (t *Engine) Delete(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echo.DELETE(path, t.wrapHandler(handler), t.handleMiddlewares(middlewares)...)

	return nil
}

// Patch 注册一个PATCH请求的路由
func (t *Engine) Patch(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echo.PATCH(path, t.wrapHandler(handler), t.handleMiddlewares(middlewares)...)

	return nil
}

// Options 注册一个OPTIONS请求的路由
func (t *Engine) Options(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echo.OPTIONS(path, t.wrapHandler(handler), t.handleMiddlewares(middlewares)...)

	return nil
}

// Any 注册一个可以处理任何HTTP方法的路由
func (t *Engine) Any(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echo.Any(path, t.wrapHandler(handler), t.handleMiddlewares(middlewares)...)

	return nil
}

// Group 创建一个新的路由组，并应用中间件
func (t *RouteGroup) Group(prefix string, middlewares ...MiddlewareFunc) *RouteGroup {

	echoGroup := t.echoGroup.Group(prefix, t.engine.handleMiddlewares(middlewares)...)

	routeGroup := &RouteGroup{
		engine:    t.engine,
		echoGroup: echoGroup,
	}

	return routeGroup
}

// Get 注册一个GET请求的路由
func (t *RouteGroup) Get(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echoGroup.GET(path, t.engine.wrapHandler(handler), t.engine.handleMiddlewares(middlewares)...)

	return nil
}

// Post 注册一个POST请求的路由
func (t *RouteGroup) Post(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echoGroup.POST(path, t.engine.wrapHandler(handler), t.engine.handleMiddlewares(middlewares)...)

	return nil
}

// Put 注册一个PUT请求的路由
func (t *RouteGroup) Put(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echoGroup.PUT(path, t.engine.wrapHandler(handler), t.engine.handleMiddlewares(middlewares)...)

	return nil
}

// Delete 注册一个DELETE请求的路由
func (t *RouteGroup) Delete(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echoGroup.DELETE(path, t.engine.wrapHandler(handler), t.engine.handleMiddlewares(middlewares)...)

	return nil
}

// Patch 注册一个PATCH请求的路由
func (t *RouteGroup) Patch(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echoGroup.PATCH(path, t.engine.wrapHandler(handler), t.engine.handleMiddlewares(middlewares)...)

	return nil
}

// Options 注册一个OPTIONS请求的路由
func (t *RouteGroup) Options(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echoGroup.OPTIONS(path, t.engine.wrapHandler(handler), t.engine.handleMiddlewares(middlewares)...)

	return nil
}

// Any 注册一个可以处理任何HTTP方法的路由
func (t *RouteGroup) Any(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) error {

	t.echoGroup.Any(path, t.engine.wrapHandler(handler), t.engine.handleMiddlewares(middlewares)...)

	return nil
}

// Start 启动HTTP服务器，监听指定的地址
func (t *Engine) Start(address string) {
	t.echo.Start(address)
}

// Shutdown 优雅地关闭HTTP服务器
// 此方法会等待正在进行的请求完成后再关闭服务器
func (t *Engine) Shutdown() error {
	return t.echo.Shutdown(context.Background())
}

// wrapHandler 将HandlerFunc转换为echo.HandlerFunc
// 该方法创建一个新的上下文并调用传入的处理函数
func (t *Engine) wrapHandler(handler HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := t.NewContext(c.Request(), c.Response().Writer)

		// 调用处理函数，如果出错则返回错误
		if err := handler(ctx); err != nil {
			return err
		}
		return nil
	}
}

// wrapMiddleware 将MiddlewareFunc转换为echo.MiddlewareFunc
// 该方法创建一个新的上下文并调用传入的中间件
func (t *Engine) wrapMiddleware(middleware MiddlewareFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := t.NewContext(c.Request(), c.Response().Writer)

			// 调用自定义中间件
			return middleware(func(ctx *Context) error {
				return next(ctx.echoContext) // 调用下一个处理函数
			})(ctx)
		}
	}
}

// handleMiddlewares 将MiddlewareFunc列表转换为echo.MiddlewareFunc列表
// 该方法用于处理传入的中间件并返回Echo兼容的中间件
func (t *Engine) handleMiddlewares(middlewares []MiddlewareFunc) []echo.MiddlewareFunc {

	var echoMiddlewares []echo.MiddlewareFunc

	// 遍历每个中间件并将其包装为echo.MiddlewareFunc
	for _, middleware := range middlewares {
		echoMiddlewares = append(echoMiddlewares, t.wrapMiddleware(middleware))
	}

	return echoMiddlewares
}
