package menu

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	controller := NewController()
	r := route.Group("menus")

	r.Get("/tree/all", controller.FindTreeAll)
	r.Get("/tree", controller.FindTree)
	r.Post("/", controller.Create)
	r.Put("/", controller.Update)
	r.Delete("/", controller.Delete)
}
