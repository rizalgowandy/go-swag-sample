package routes

import (
	"github.com/labstack/echo/v4"
	_ "github.com/rizalgowandy/go-swag-sample/docs/echomicroservice/app1"
	"github.com/rizalgowandy/go-swag-sample/echomicroservice/internal/app1/controller"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Swagger Example API
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
func Register(e *echo.Echo) {
	e.GET("/", controller.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
