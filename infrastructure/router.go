package infrastructure

import (
	"github.com/labstack/echo/v4"

	"github.com/Luftalian/writers_app/interfaces/controller"
)

var Router *echo.Echo

func init() {
	e := echo.New()

	userController := controller.NewUserController(NewDbHandler())

	e.POST("/users", func(c echo.Context) error {
		userController.Create(c)
		return nil
	})
	e.GET("/users", func(c echo.Context) error {
		userController.Index(c)
		return nil
	})
	e.GET("/users/:id", func(c echo.Context) error {
		userController.Show(c)
		return nil
	})

	Router = e
}
