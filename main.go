package main

import (
	"breeze-api/config"
	"breeze-api/pkg/dal"
	"breeze-api/internal/router"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// 实例化
	app := fiber.New()

	// 初始化数据访问层
	dal.InitDal(&dal.Config{})

	// 根目录
	app.Static("/", "./web")

	// 注册路由
	router.ApiRegister(app)

	app.Listen(config.App.Host + ":" + strconv.Itoa(config.App.Port))
}
