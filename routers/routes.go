package routers

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	// Users
	e.POST("/login", controllers.LoginUser)
	e.POST("/users", controllers.CreateUserControllers)
	e.GET("/users", controllers.GetAllusercontrollers)
	e.GET("/users/:uid", controllers.GetUserControllers)
	e.PUT("/users/:uid", controllers.UpdateUserControllers)
	e.DELETE("/users/:uid", controllers.DeleteUserControllers)

	// Products
	e.POST("/products", controllers.CreateProductControllers)
	e.GET("/products", controllers.GetAllProductControllers)
	e.GET("/products/:pid", controllers.GetProductControllers)
	e.PUT("/products/:pid", controllers.UpdateProductControllers)
	e.DELETE("/products/:pid", controllers.DeleteProductControllers)

	// Spesific Products
	e.GET("/products/PaketData", controllers.GetPaketData)
	e.GET("/products/Pulsa", controllers.GetPulsa)
	e.GET("/products/Emoney", controllers.GetEmoney)
	e.GET("/products/Cashout", controllers.GetCashout)

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
