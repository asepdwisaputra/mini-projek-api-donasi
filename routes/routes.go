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

	// User
	e.POST("/users", controllers.CreateUserController)
	// e.GET("/users", controllers.GetUsersController)
	// e.GET("/users/:id", controllers.GetUserController)
	// e.DELETE("/users/:id", controllers.DeleteUserController)
	// e.PUT("/users/:id", controllers.UpdateUserController)

	// Login User
	e.POST("/login", controllers.LoginUserController) // Create token

	// Campaign
	e.GET("/campaigns", controllers.GetCampaigns)
	e.GET("/campaigns/:id", controllers.GetCampaign)
	//e.POST("/campaigns", controllers.CreateCampaign)

	// Donation
	// e.GET("/donations", controllers.GetDonations)
	// e.GET("/donations/:id", controllers.GetDonationByID)
	// e.GET("/donations/user/:id", controllers.GetDonationsByUserID)
	// e.POST("/donations", controllers.CreateDonation)

	//Logger Middleware
	middleware.LogMiddleware(e)
	// Akan Muncul di Console --> method=GET, uri=/users, status=200, latency_human=534.7µs

	//Basic Auth Databse
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	//eAuthBasic.GET("/users", controllers.GetUsersController)

	//JWT
	eJWT := e.Group("/jwt")
	eJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	// User
	eJWT.GET("/users", controllers.GetUsersController)
	eJWT.GET("/users/:id", controllers.GetUserController)
	eJWT.PUT("/users/:id", controllers.UpdateUserController)
	eJWT.DELETE("/users/:id", controllers.DeleteUserController)

	// Campaign
	eJWT.POST("/campaigns", controllers.CreateCampaign)

	// Donation
	eJWT.GET("/donations", controllers.GetDonations)
	eJWT.GET("/donations/:id", controllers.GetDonationByID)
	eJWT.GET("/donations/user/:id", controllers.GetDonationsByUserID)
	eJWT.POST("/donations", controllers.CreateDonation)

	return e
}
