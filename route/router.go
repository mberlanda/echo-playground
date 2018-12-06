package route

import (
	"net/http"

	"github.com/mberlanda/echo-playground/api"

	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Debug = true

	// Middlewares
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := e.Group("/api/v1")
	v1.POST("/users", api.CreateUser)
	v1.GET("/users/:id", api.GetUser)
	v1.PUT("/users/:id", api.UpdateUser)
	v1.DELETE("/users/:id", api.DeleteUser)

	return e
}
