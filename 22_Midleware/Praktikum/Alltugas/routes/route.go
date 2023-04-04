package routes

import (
	"go_Muhammad-Wahyu-Yudiansyah/22_Midleware/Praktikum/Alltugas/constant"
	"go_Muhammad-Wahyu-Yudiansyah/22_Midleware/Praktikum/Alltugas/controller"
	"go_Muhammad-Wahyu-Yudiansyah/22_Midleware/Praktikum/Alltugas/midleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	midleware.LogMiddleware(e)
	// routing with query parameter

	e.POST("/login", controller.LoginUserController)
	// routing with query parameter
	e.GET("/blogs", controller.GetBlogsController)
	e.POST("/blogs", controller.CreateBlogController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(middleware.BasicAuth(midleware.BasicAuthDb))
	eAuthBasic.GET("/users", controller.GetUsersController)
	eAuthBasic.GET("/users/:id", controller.GetUserController)
	eAuthBasic.DELETE("/users/:id", controller.DeleteUserController)
	eAuthBasic.PUT("/users/:id", controller.UpdateUserController)
	eAuthBasic.GET("/books", controller.GetBooksController)
	eAuthBasic.POST("/books", controller.CreateBookController)
	eAuthBasic.GET("/books/:id", controller.GetBookController)
	eAuthBasic.DELETE("/books/:id", controller.DeleteBookController)
	eAuthBasic.PUT("/books/:id", controller.UpdateBookController)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("/users", controller.GetUsersController)
	eJwt.GET("/users/:id", controller.GetUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)
	eJwt.GET("/books", controller.GetBooksController)
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.GET("/books/:id", controller.GetBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
	return e
}
