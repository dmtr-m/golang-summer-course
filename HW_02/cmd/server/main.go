package main

import (
	"httpProject/accounts"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccountDetails)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.POST("/account/delete", accountsHandler.DeleteAccount)
	e.POST("/account/changebalance", accountsHandler.ChangeAccountBalance)
	e.POST("/account/changename", accountsHandler.ChangeAccountName)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
