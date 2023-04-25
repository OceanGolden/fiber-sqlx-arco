package main

import (
	"fiber-sqlx-arco/pkg/configs"
	"fiber-sqlx-arco/pkg/global"
	"fiber-sqlx-arco/pkg/middlewares"
	"fiber-sqlx-arco/pkg/routes"
	"fiber-sqlx-arco/pkg/utils"
	"fiber-sqlx-arco/platform/database"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	// 初始化
	configs.Init()
	// 连接数据库
	global.DB = database.OpenDBConnection()
	// fiber 自身配置
	fiberConfig := configs.FiberConfig()
	// 创建实例
	app := fiber.New(fiberConfig)
	// 中间件
	middlewares.FiberMiddleware(app) // 注册 fiber内置中间件
	// 路由.
	//routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	// routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
