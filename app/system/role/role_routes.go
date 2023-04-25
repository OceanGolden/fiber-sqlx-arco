package role

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	controller := NewController()
	r := route.Group("roles")

	r.Get("/", controller.FindPage)
	r.Post("/", controller.Create)
	r.Put("/", controller.Update)
	r.Delete("/", controller.Delete)
	r.Get("/all", controller.FindAll)
	r.Get("/menus", controller.FindMenus)
	r.Post("/menus", controller.GrantMenu)

}
