package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"integra.com/go/cmd/handlers"
	"integra.com/go/cmd/storage"
)

func main() {
	app := echo.New()          // Create a new echo instance
	app.Use(middleware.CORS()) // Enable CORS

	port := ":1323"  // Port number
	version := "/v1" // API version

	storage.InitDB() // Initialize the database

	app.GET(version+"/user", handlers.GetUsers)
	app.POST(version+"/user", handlers.CreateUser)
	app.PUT(version+"/user/:id", handlers.UpdateUser)
	app.DELETE(version+"/user/:id", handlers.DeleteUser)

	app.Logger.Fatal(app.Start(port))
}
