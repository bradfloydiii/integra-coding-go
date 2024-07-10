package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
)

func createUser(c echo.Context) error {
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	email := c.FormValue("email")
	company := c.FormValue("company")
	phone := c.FormValue("phone")
	return c.String(http.StatusOK, "firstName: "+firstName+", lastName: "+lastName+", email: "+email+", company: "+company+", phone: "+phone)
}

type User struct {
	ID        int    `json:"_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Company   string `json:"company"`
	Phone     string `json:"phone"`
}

func getUsers(c echo.Context) error {
	users := []User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Company: "ABC Inc.", Phone: "1234567890"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Company: "XYZ Corp.", Phone: "9876543210"},
		// Add more users as needed
	}

	return c.JSON(http.StatusOK, users)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	email := c.FormValue("email")
	company := c.FormValue("company")
	phone := c.FormValue("phone")
	return c.String(http.StatusOK, "id: "+id+", firstName: "+firstName+", lastName: "+lastName+", email: "+email+", company: "+company+", phone: "+phone)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Deletes user in database with id="+id)
}

func main() {
	app := echo.New() // Create a new echo instance

	app.Use(middleware.CORS()) // Enable CORS

	port := ":1323"  // Port number
	version := "/v1" // API version

	// Routes
	app.POST(version+"/user", createUser)
	app.GET(version+"/user", getUsers)
	app.PUT(version+"/user/:id", updateUser)
	app.DELETE(version+"/user/:id", deleteUser)

	// Start server
	app.Logger.Fatal(app.Start(port))
}
