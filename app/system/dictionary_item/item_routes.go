package dictionary_item

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	controller := NewController()
	dictionaryRoutes := route.Group("dictionary")
	r := dictionaryRoutes.Group("items")

	r.Get("/", controller.FindPage)
	r.Post("/", controller.Create)
	r.Put("/", controller.Update)
	r.Delete("/", controller.Delete)
}
