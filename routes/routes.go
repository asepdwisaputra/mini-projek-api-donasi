package routes

import (
	"api-donasi/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/user/:id", controllers.GetUser)
	e.POST("/campaign", controllers.CreateCampaign)
	e.POST("/donation", controllers.MakeDonation)
	// ... tambahkan rute lain sesuai kebutuhan Anda
}
