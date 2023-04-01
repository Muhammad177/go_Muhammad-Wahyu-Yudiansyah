package routes

import (
	"go_Muhammad-Wahyu-Yudiansyah/21_ORM-MVC/Praktikum/Alltugas/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// routing with query parameter
	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/users/:id", controller.GetUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	// routing with query parameter
	e.GET("/books", controller.GetBooksController)
	e.POST("/books", controller.CreateBookController)
	e.GET("/books/:id", controller.GetBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	e.PUT("/books/:id", controller.UpdateBookController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
	return e
}
