package main

import (
	"fmt"
	"go-license/internal/config"
	"go-license/internal/http/handlers/student"
	"go-license/internal/storage/sqllite"
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

// type application struct {
// 	config config
// }

// type config struct {
// 	addr string
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
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// database setup

	_,err :=sqllite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}


	slog.Info("storage initialised")

	// setup router

	r := fiber.New()

	r.Post("/api/students", student.New)

	r.Listen(cfg.Port)

}
