package routes

import (
	"github.com/emanuelhardwell/go-user/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	g := app.Group("/api/users")
	g.Post("/", controller.Create)
}
