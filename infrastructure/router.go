package infrastructure

import (
	"github.com/labstack/echo/v4"

	"github.com/Luftalian/writers_app/interfaces/controller"
)

var Router *echo.Echo

func init() {
	e := echo.New()

	dbHandler := NewDbHandler()

	userController := controller.NewUserController(dbHandler)

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

	textController := controller.NewTextController(dbHandler)
	e.POST("/texts", func(c echo.Context) error {
		textController.Create(c)
		return nil
	})
	e.GET("/texts", func(c echo.Context) error {
		textController.Index(c)
		return nil
	})
	e.GET("/texts/:id", func(c echo.Context) error {
		textController.Show(c)
		return nil
	})
	e.PUT("/texts/:id", func(c echo.Context) error {
		textController.Change(c)
		return nil
	})
	e.DELETE("/texts/:id", func(c echo.Context) error {
		textController.Delete(c)
		return nil
	})

	tagController := controller.NewTagController(dbHandler)
	e.POST("/tags", func(c echo.Context) error {
		tagController.Create(c)
		return nil
	})
	e.GET("/tags", func(c echo.Context) error {
		tagController.Index(c)
		return nil
	})
	e.GET("/tags/:id", func(c echo.Context) error {
		tagController.Show(c)
		return nil
	})

	Router = e
}
