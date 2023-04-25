package auth

import (
	"fiber-sqlx-arco/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(route fiber.Router) {
	controller := NewController()
	r := route.Group("auth")

	r.Post("/login", controller.Login)
	r.Post("/logout", middlewares.JWTProtected(), controller.Logout)
	r.Get("/info", middlewares.JWTProtected(), controller.FindInfo)
	//r.Delete("/", handler.Delete)
	//r.Post("/info", controller.Info)

}
