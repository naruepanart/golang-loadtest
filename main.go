package main

import (
	"net/http"

	"github.com/labstack/echo"

)

func getUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware

	// Route => handler
	e.GET("/", getUsers)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
