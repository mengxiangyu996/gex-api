# Gex API

> Gex API 是一个构建在 Go 语言之上的轻量级 Web 框架，它以 Echo 框架为核心，提供了一套高效、灵活的 API 设计和路由管理功能。该框架致力于简化开发过程，同时确保系统的稳定性和可扩展性。

### 使用说明
1. ###### 克隆项目
```
  git clone https://github.com/mengxiangyu996/gex-api.git
```
2. ###### 进入目录
```
  cd gex-api
```
3. ###### 修改配置文件
```
  mv env.json.example env.json
```
4. ###### 安装依赖
```
  go mod tidy
```
5. ###### 启动服务
```
  go run cmd\main.go
```

### 阿里云流水线
- #### 构建
- ###### 对于常规构建，可以使用以下命令
```
    # default use of goproxy.cn
    export GOPROXY=https://goproxy.cn
    # input your command here
    go build main.go
```
- ###### 如果需要为特定平台（例如 arm64 架构的 Linux）构建，可以使用以下命令
```
    export GOPROXY=https://goproxy.cn
    # input your command here
    GOARCH=arm64 GOOS=linux go build -o main main.go
```
- ###### 构建物上传配置
- ![image](http://fiber-api.ddnsgeek.com/20240927103036.png)


- #### 部署
- ###### 部署命令
```
    tar zxvf 
    /opt/1panel/apps/openresty/openresty/www/sites/breeze/package.tgz -C 
    /opt/1panel/apps/openresty/openresty/www/sites/breeze/
    cd /opt/1panel/apps/openresty/openresty/www/sites/breeze
    bash service.sh restart
```

### 技术栈

- **编程语言**: [Go](https://go.dev) 一种简洁、高效的编程语言，特别适用于构建高性能的网络服务。
- **Web 框架**: [Echo](https://echo.labstack.com) 一个轻量级的 Web 框架，以其高性能和灵活的路由系统著称。
- **ORM**: [Gorm](https://gorm.io) 一个功能丰富的 ORM，简化了数据库操作。
- **缓存**: [Redis](https://redis.uptrace.dev) 一个高性能的键值存储系统，用于缓存和会话管理等。

### 特性

  - **简单易用**: 直观的设计，快速上手。
  - **灵活的路由管理**: 支持路由分组和中间件应用。
  - **高性能**: 基于 Echo 框架，提供高效的服务。

- ####  标签默认值功能

> - Gex API 在 `Context` 中增加了 `BindX` 方法，这是对传统 `Bind` 方法的扩展。它允许开发者通过 `tag` 为请求参数设置默认值，从而简化数据绑定和验证过程。
```go
    var param struct {
        Page      int    `default:"1"`
        Size      int    `default:"10"`
        IsBool    bool   `default:"true"`
        Keyword   string `default:"hello world"`
        Number    float64 `default:"1.23"`
        Status    [3]int `default:"1,2"`
        TimeRange []string `default:"12:30,13:30"`
        Data      map[string]interface{} `default:"age:18,name:zero"`
    }

    if err := ctx.BindX(&param); err != nil {
        return ctx.Fail(err.Error())
    }

```
### 目录结构
```
├─app
│  ├─controller			(控制器层)
│  ├─database			(数据库相关的脚本和迁移文件)
│  ├─internal			(内部模块，包含一些核心功能)
│  │  ├─encrypt			(加密相关功能)
│  │  ├─jwt			(JWT 令牌相关功能)
│  │  ├─request			(请求工具)
│  │  └─utils			(工具函数，提供一些通用的功能)
│  ├─middleware			(中间件)
│  ├─model			(数据模型，定义数据结构和数据库交互)
│  ├─request			(请求结构体，定义 API 请求的参数)
│  ├─response			(响应结构体，定义 API 返回的数据格式)
│  ├─route			(路由定义，配置 API 路由)
│  └─service			(业务逻辑层，处理具体的业务逻辑)
├─cmd				(主程序入口，启动应用的相关代码)
├─config			(配置文件，存放应用的配置项)
└─pkg				(包含公共库和工具包)
    ├─builder			(封装的引擎和上下文，构建服务的基础)
    ├─dal			(数据访问层，封装与数据库的交互)
    ├─datetime			(日期时间处理相关功能)
    ├─env			(配置读取)
    ├─middleware		(公共中间件)
    └─tag			(标签功能相关实现)
```
