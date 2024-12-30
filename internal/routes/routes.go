package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamrwtktw/lic/internal/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/users", handlers.CreateUser)

}
