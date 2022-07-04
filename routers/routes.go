package routers

import (
	"backend/constants"
	"backend/controllers"

	// "backend/middleware"

	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	//login
	e.POST("/user/login", controllers.LoginUserController)
	e.POST("/admin/login", controllers.LoginAdminController)

	// Users
	e.POST("/users", controllers.CreateUserControllers)
	e.GET("/users", controllers.GetAllusercontrollers, eMiddleware.JWT([]byte(constants.SECRET_JWT)))
	e.GET("/users/:uid", controllers.GetUserControllers)
	e.PUT("/users/:uid", controllers.UpdateUserControllers)
	e.DELETE("/users/:uid", controllers.DeleteUserControllers)

	// Products
	e.POST("/products", controllers.CreateProductControllers)
	e.GET("/products", controllers.GetAllProductControllers, eMiddleware.JWT([]byte(constants.SECRET_JWT)))
	e.GET("/products/:pid", controllers.GetProductControllers)
	e.PUT("/products/:pid", controllers.UpdateProductControllers)
	e.DELETE("/products/:pid", controllers.DeleteProductControllers)

	// Spesific Products
	e.GET("/products/PaketData", controllers.GetPaketData)
	e.GET("/products/Pulsa", controllers.GetPulsa)
	e.GET("/products/Emoney", controllers.GetEmoney)
	e.GET("/products/Cashout", controllers.GetCashout)

	//transaction
	e.GET("/redeem", controllers.GetAllRedeemControllers)
	e.GET("/user/redeem/:user_id", controllers.GetRedeemByUserIDControllers)
	e.POST("/redeem", controllers.CreateRedeemsControllers)
	e.PUT("/redeem/:id", controllers.UpdateRedeemControllers, eMiddleware.JWT([]byte(constants.SECRET_JWT)))
	e.DELETE("/redeem/:id", controllers.DeleteRedeemControllers)

	//transaction
	e.GET("/transaction", controllers.GetAllTransactionControllers)
	e.GET("/transaction/user/:id", controllers.GetTransactionByIdUserControllers)
	e.GET("/transaction/:id", controllers.GetTransactionByIdControllers)
	e.POST("/transaction", controllers.CreateTransactionsControllers)
	e.PUT("/transaction", controllers.UpdatetransactionControllers)
	e.DELETE("/transaction", controllers.DeleteTansactionControllers)

	return e
}
