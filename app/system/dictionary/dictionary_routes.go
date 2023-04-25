package dictionary

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	controller := NewController()
	r := route.Group("dictionaries")

	r.Get("/items", controller.FindItems)
	r.Get("/", controller.FindPage)
	r.Post("/", controller.Create)
	r.Put("/", controller.Update)
	r.Delete("/", controller.Delete)
}
