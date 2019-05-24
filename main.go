package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	username string `bson:"username"`
	password string `bson:"password"`
}

func getUsers(c echo.Context) error {
	return c.String(http.StatusOK, "getUsers\n")
}
func createUsers(c echo.Context) error {
	return nil
}
func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "deleteUser\n")
}

const (
	mongoURL = "mongodb+srv://guload:mNW4RK1skR2yTdT5@guload-singapore-w5sv8.gcp.mongodb.net/main?retryWrites=true"
)

func main() {
	e := echo.New()
	
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Route => handler
	e.GET("/", getUsers)
	e.POST("/", createUsers)
	e.DELETE("/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
