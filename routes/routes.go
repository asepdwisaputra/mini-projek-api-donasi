package routes

import (
	"api-donasi/constants"
	"api-donasi/controllers"
	"api-donasi/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.POST("/users/login", controllers.LoginUserController)

	e.GET("/campaigns", controllers.GetCampaigns)
	e.POST("/campaigns", controllers.CreateCampaign)
	e.POST("/donation", controllers.CreateDonation)

	//Logger Middleware
	middleware.LogMiddleware(e)
	// Akan Muncul di Console --> method=GET, uri=/users, status=200, latency_human=534.7µs

	//Basic Auth Databse
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/users", controllers.GetUserController)

	//JWT
	eJWT := e.Group("/jwt")
	eJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJWT.GET("/users", controllers.GetUsersController)
	eJWT.GET("/users/:id", controllers.GetUserController)
	eJWT.PUT("/users/:id", controllers.UpdateUserController)
	eJWT.DELETE("/users/:id", controllers.DeleteUserController)

	return e
}