package routers

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/user", controllers.GetAllusercontrollers)
	e.GET("/user/:id", controllers.Getusercontrollers)
	e.PUT("/user/update/:id", controllers.Updateusercontrollers)
	e.POST("/user", controllers.Createusercontrollers)
	e.DELETE("/user/:id", controllers.Deleteusercontrollers)

	return e
}
