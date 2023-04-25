package staff

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	controller := NewController()
	r := route.Group("staffs")

	r.Get("/", controller.FindPage)
	r.Post("/", controller.Create)
	r.Put("/", controller.Update)
	r.Delete("/", controller.Delete)
	r.Post("/roles", controller.AssignRole)
	//r.Get("/info", controller.FindInfo)
	//staff.Get("/info", handler.Info)
	//staff.Get("/role/:id", handler.FindRole)
	//staff.Post("/assign", handler.Assign)
}
