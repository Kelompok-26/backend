package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("api/users/", controllers.CreateUser)
	e.GET("api/users/", controllers.GetAllUsers)
	e.GET("api/users/:UID", controllers.GetSpesificUser)
	e.PUT("/user/update/:id", controllers.EDITUser)
	e.DELETE("api/users/:UID", controllers.DeleteUser)

	return e
}
