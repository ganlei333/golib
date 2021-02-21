package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//App 代理http接口对象
var App *fiber.App

//InitHTTPServer 注册路由
func InitHTTPServer(log bool, port string) {
	App = fiber.New()
	if log {
		config := logger.Config{
			Format:     "${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
			TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
		}
		App.Use(logger.New(config))

	}
	App.Get("/ping", ping)
	App.Listen(":" + port)
}
