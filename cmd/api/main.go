package main

import (
	"go-license/internal/config"

	"github.com/gofiber/fiber/v2"
)

// type application struct {
// 	config config
// }

// type config struct {
// 	addr string
// }

// func (app *application) healthCheckHandler(c *fiber.Ctx) error {
// 	return c.JSON(fiber.Map{
// 		"status":  "success",
// 		"message": "Health check passed",
// 	})
// }

// func (app *application) mount() *fiber.App {

// 	//  port := os.Getenv("PORT")

// 	//  fmt.Println(port)

// 	r := fiber.New()

// 	// Define routes
// 	r.Get("/", app.healthCheckHandler)

// 	// Start the server
// 	r.Listen(":3000")

// 	return r
// }


func main() {

	// load config
	cfg :=config.MustLoad()
	

	// setup router

	r := fiber.New()


	
	r.Listen(cfg.Port)



}
