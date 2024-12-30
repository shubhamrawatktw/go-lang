package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamrwtktw/lic/internal/config/env"
	"github.com/shubhamrwtktw/lic/internal/routes"
)


func InitApp(){
	
	
	app:= fiber.New()

	routes.SetupRoutes(app)
	app.Listen(env.Env.PORT)
}