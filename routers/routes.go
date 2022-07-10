package routers

import (
	"backend/constants"
	"backend/controllers"
	"backend/middleware"

	// "backend/middleware"

	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Use(eMiddleware.CORS())
	eJwt := e.Group("/v1")
	eJwt.Use(eMiddleware.JWT([]byte(constants.SECRET_JWT)))

	// Login
	e.POST("/user/login", controllers.LoginUserController)
	e.POST("/admin/login", controllers.LoginAdminController)

	// Users
	e.POST("/users", controllers.CreateUserControllers)
	eJwt.GET("/users", controllers.GetAllusercontrollers, middleware.AdminRole)
	eJwt.GET("/users/:uid", controllers.GetUserControllers, middleware.AdminRoleorUserID)
	eJwt.PUT("/users/update/:uid", controllers.UpdateUserControllers)
	eJwt.PUT("/addpointusers/:id", controllers.AddPointUserController, middleware.AdminRole)
	eJwt.DELETE("/users/:uid", controllers.DeleteUserControllers, middleware.AdminRole)

	// Products
	eJwt.POST("/products", controllers.CreateProductControllers)
	e.GET("/products", controllers.GetAllProductControllers)
	e.GET("/products/:pid", controllers.GetProductControllers)
	eJwt.PUT("/products/update/:pid", controllers.UpdateProductControllers, middleware.AdminRole)
	eJwt.DELETE("/products/:pid", controllers.DeleteProductControllers, middleware.AdminRole)

	// Spesific Products
	e.GET("/products/PaketData", controllers.GetPaketData)
	e.GET("/products/Pulsa", controllers.GetPulsa)
	e.GET("/products/Emoney", controllers.GetEmoney)
	e.GET("/products/Cashout", controllers.GetCashout)

	// Redeemm
	eJwt.GET("/redeem", controllers.GetAllRedeemControllers, middleware.AdminRole)
	eJwt.GET("/user/redeem/:user_id", controllers.GetRedeemByUserIDControllers)
	eJwt.POST("/redeem", controllers.CreateRedeemsControllers)
	eJwt.PUT("/redeem/:id", controllers.UpdateRedeemControllers, middleware.AdminRole)
	eJwt.DELETE("/redeem/:id", controllers.DeleteRedeemControllers, middleware.AdminRole)

	// Transaction
	eJwt.GET("/transaction", controllers.GetAllTransactionControllers, middleware.AdminRole)
	eJwt.GET("/user/transaction/:user_id", controllers.GetTransactionByIdUserControllers)
	eJwt.GET("/transaction/:id", controllers.GetTransactionByIdControllers)
	eJwt.POST("/transaction", controllers.CreateTransactionsControllers)
	eJwt.PUT("/transaction/:id", controllers.UpdatetransactionControllers, middleware.AdminRole)
	eJwt.DELETE("/transaction/:id", controllers.DeleteTansactionControllers, middleware.AdminRole)

	return e
}
