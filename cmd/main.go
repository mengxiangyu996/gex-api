package main

import (
	"gex-api/config"
	"gex-api/database"
	"gex-api/pkg/builder"
	"gex-api/pkg/dal"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := config.Mysql.Username + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Host + ":" + config.Mysql.Port + ")/" + config.Mysql.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	// 传入配置用来链接mysql和redis
	// 以 server := builder.New(&builder.Config{}) 初始化的时候不会初始化mysql服务，需注释掉下方初始化数据库的流程
	
	server := builder.New(&builder.Config{
		GormConfig: &dal.GormConfig{
			Dialector: mysql.Open(dsn),
			Opts: &gorm.Config{
				SkipDefaultTransaction: true, // 跳过默认事务
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
				Logger: logger.New(log.Default(), logger.Config{
					LogLevel: logger.Silent, // 不打印日志
				}),
			},
		},
	})

	// 初始化数据库
	database.Init()

	// 配置静态文件
	server.Static("/", "web")

	server.Get("/hello", func(ctx *builder.Context) error {
		return ctx.String(200, "Hello, World!")
	})

	// 启动服务
	server.Start(config.App.Host)
}
