package main

import (
	"log"
	"net/http"

	swagger "github.com/arsmn/fiber-swagger"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/rizalgowandy/go-swag-sample/docs/fibersimple"
)

// @title Fiber Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	// Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(middleware.Recover())

	// Routes
	app.Get("/", HealthCheck)
	app.Use("/swagger", swagger.Handler) // default

	// 404 Handler
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(http.StatusNotFound) // => 404 "Not Found"
	})

	// Start Server
	if err := app.Listen(3000); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		c.Next(err)
	}
}
