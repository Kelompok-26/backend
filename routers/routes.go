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

	v1 := e.Group("/v1")

	eJwt := v1.Group("")
	
	

	eJwt.Use(eMiddleware.JWT([]byte(constants.SECRET_JWT)))

	// Login
	v1.POST("/user/login", controllers.LoginUserController)
	v1.POST("/admin/login", controllers.LoginAdminController)

	// Users
	v1.POST("/users", controllers.CreateUserControllers)
	eJwt.GET("/users", controllers.GetAllusercontrollers, middleware.AdminRole)
	eJwt.GET("/users/:uid", controllers.GetUserControllers, middleware.AdminRoleorUserID)
	eJwt.PUT("/users/update/:uid", controllers.UpdateUserControllers)
	eJwt.PUT("/addpointusers/:id", controllers.AddPointUserController, middleware.AdminRole)
	eJwt.DELETE("/users/:uid", controllers.DeleteUserControllers, middleware.AdminRole)

	// Products
	eJwt.POST("/products", controllers.CreateProductControllers)
	v1.GET("/products", controllers.GetAllProductControllers)
	v1.GET("/products/:pid", controllers.GetProductControllers)
	eJwt.PUT("/products/update/:pid", controllers.UpdateProductControllers, middleware.AdminRole)
	eJwt.DELETE("/products/:pid", controllers.DeleteProductControllers, middleware.AdminRole)

	// Spesific Products
	v1.GET("/products/PaketData", controllers.GetPaketData)
	v1.GET("/products/Pulsa", controllers.GetPulsa)
	v1.GET("/products/Emoney", controllers.GetEmoney)
	v1.GET("/products/Cashout", controllers.GetCashout)

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
