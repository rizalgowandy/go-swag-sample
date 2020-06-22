package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizalgowandy/go-swag-sample/echomicroservice/internal/app1/routes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	routes.Register(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
