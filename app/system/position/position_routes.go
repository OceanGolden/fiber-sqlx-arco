package position

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	controller := NewController()
	r := route.Group("positions")

	r.Get("/all", controller.FindAll)
	r.Get("/", controller.FindPage)
	r.Post("/", controller.Create)
	r.Put("/", controller.Update)
	r.Delete("/", controller.Delete)
}
