package routers

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/user", controllers.GetAllusercontrollers)
	e.GET("/user/:id", controllers.GetUserControllers)
	e.PUT("/user/:id", controllers.UpdateUserControllers)
	e.POST("/user", controllers.CreateUserControllers)
	e.DELETE("/user/:id", controllers.DeleteUserControllers)

	e.GET("/product", controllers.GetAllProductControllers)
	e.GET("/product/:id", controllers.GetProductControllers)
	e.PUT("/product/:id", controllers.UpdateProductControllers)
	e.POST("/product", controllers.CreateProductControllers)
	e.DELETE("/product/:id", controllers.DeleteProductControllers)

	e.GET("/redeem", controllers.GetAllRedeemControllers)
	e.GET("/user/redeem/:user_id", controllers.GetRedeemByUserIDControllers)
	e.POST("/redeem", controllers.CreateRedeemsControllers)
	e.PUT("/redeem/:id", controllers.UpdateRedeemControllers)
	e.DELETE("/redeem/:id", controllers.DeleteRedeemControllers)

	e.GET("/transaction", controllers.GetAllTransactionControllers)
	e.GET("/transaction/:id", controllers.GetTransactionIDControllers)
	e.POST("/transaction", controllers.CreateTransactionsControllers)

	return e
}
