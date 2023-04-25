package routes

import (
	"fiber-sqlx-arco/app/system/dictionary"
	"fiber-sqlx-arco/app/system/dictionary_item"
	"fiber-sqlx-arco/app/system/menu"
	"fiber-sqlx-arco/app/system/organization"
	"fiber-sqlx-arco/app/system/position"
	"fiber-sqlx-arco/app/system/role"
	"fiber-sqlx-arco/app/system/staff"
	"fiber-sqlx-arco/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	appRoute := a.Group("/")
	//system := appRoute.Group("/system")
	system := appRoute.Group("/system", middlewares.JWTProtected(), middlewares.CasbinProtected())
	//system := appRoute.Group("/system", middlewares.JWTProtected())

	staff.InitRoutes(system)
	role.InitRoutes(system)
	position.InitRoutes(system)
	organization.InitRoutes(system)
	dictionary.InitRoutes(system)
	dictionary_item.InitRoutes(system)
	menu.InitRoutes(system)
}
