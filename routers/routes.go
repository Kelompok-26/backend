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
	eJwt.GET("/users/:id", controllers.GetUserControllers, middleware.AdminRoleorUserID)
	eJwt.PUT("/users/update/:uid", controllers.UpdateUserControllers, middleware.AdminRoleorUserID)
	eJwt.PUT("/users/:id/point", controllers.AddPointUserController, middleware.AdminRole)
	eJwt.DELETE("/users/:uid", controllers.DeleteUserControllers, middleware.AdminRole)

	// Products
	eJwt.POST("/products", controllers.CreateProductControllers, middleware.AdminRole)
	v1.GET("/products", controllers.GetAllProductControllers)
	v1.GET("/products/:pid", controllers.GetProductControllers)
	eJwt.PUT("/products/update/:id", controllers.UpdateProductControllers, middleware.AdminRole)
	eJwt.DELETE("/products/:pid", controllers.DeleteProductControllers, middleware.AdminRole)

	// Spesific Products
	v1.GET("/products/PaketData", controllers.GetPaketData)
	v1.GET("/products/Pulsa", controllers.GetPulsa)
	v1.GET("/products/E-Money", controllers.GetEmoney)
	v1.GET("/products/Cashout", controllers.GetCashout)

	// Transaction
	eJwt.GET("/transaction", controllers.GetAllTransactionControllers, middleware.AdminRole)
	eJwt.GET("/user/transaction/:user_id", controllers.GetTransactionByIdUserControllers, middleware.AdminRoleorUserID)
	eJwt.GET("/transaction/:id", controllers.GetTransactionByIdControllers, middleware.AdminRole)
	eJwt.POST("/user/:id/transaction", controllers.UserCreateTransactionsController, middleware.AdminRoleorUserID)
	eJwt.DELETE("/transaction/:id", controllers.DeleteTansactionControllers, middleware.AdminRole)

	return e
}
