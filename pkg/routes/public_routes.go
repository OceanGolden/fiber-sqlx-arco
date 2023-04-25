package routes

import (
	"fiber-sqlx-arco/app/security/auth"
	"fiber-sqlx-arco/pkg/common/response"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/")
	route.Get("xx", func(ctx *fiber.Ctx) error {
		return ctx.JSON(response.OK("sxx"))
	})
	auth.InitRoutes(route)
}
