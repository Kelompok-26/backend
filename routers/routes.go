package routers

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/user", controllers.GetAllusercontrollers)
	e.GET("/user/:id", controllers.Getusercontrollers)
	e.PUT("/user/:id", controllers.Updateusercontrollers)
	e.POST("/user", controllers.Createusercontrollers)
	e.DELETE("/user/:id", controllers.Deleteusercontrollers)

	e.GET("/product", controllers.GetAllproductcontrollers)
	e.GET("/product/:id", controllers.Getproductcontrollers)
	e.PUT("/product/:id", controllers.Updateproductcontrollers)
	e.POST("/product", controllers.Createproductcontrollers)
	e.DELETE("/product/:id", controllers.Deleteproductcontrollers)

	e.GET("/redeem", controllers.GetAllredeemcontrollers)
	e.GET("/user/redeem/:user_id", controllers.GetredeemByUserIDcontrollers)
	e.POST("/redeem", controllers.CreateRedeemscontrollers)

	return e
}
