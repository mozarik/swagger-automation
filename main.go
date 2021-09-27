package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pdrum/swagger-automation/api"

	_ "github.com/pdrum/swagger-automation/docs" // This line is necessary for go-swagger to find your docs!
)

func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Set up basic auth with username=foo and password=bar
	// e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
	// 	Validator: func(username, password string, c echo.Context) (bool, error) {
	// 		if username == "foo" && password == "bar" {
	// 			return true, nil
	// 		}
	// 		return false, nil
	// 	},
	// }))

	// Route => handler
	e.POST("/foobar", api.FooBarHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
