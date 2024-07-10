package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User represents a user in the system.
type User struct {
	ID        int    `json:"_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Company   string `json:"company"`
	Phone     string `json:"phone"`
}

// swagger:route GET /v1/user users getUsers
// Get a list of users.
// responses:
//   200: usersResponse

// getUsers handles the GET request for retrieving a list of users.
func getUsers(c echo.Context) error {
	// users := sq.Select("*").From("users").Join("emails USING (email_id)")

	users := []User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Company: "ABC Inc.", Phone: "1234567890"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Company: "XYZ Corp.", Phone: "9876543210"},
		// Add more users as needed
	}

	return c.JSON(http.StatusOK, users)
}

// swagger:route POST /v1/user users createUser
// Create a new user.
// responses:
//   200: userResponse

// createUser handles the POST request for creating a new user.
func createUser(c echo.Context) error {
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	email := c.FormValue("email")
	company := c.FormValue("company")
	phone := c.FormValue("phone")

	fmt.Println(http.StatusOK, "firstName: "+firstName+", lastName: "+lastName+", email: "+email+", company: "+company+", phone: "+phone)

	return c.JSON(http.StatusOK, []User{})
}

// swagger:route PUT /v1/user/{id} users updateUser
// Update an existing user.
// responses:
//   200: userResponse

// updateUser handles the PUT request for updating an existing user.
func updateUser(c echo.Context) error {
	id := c.Param("id")
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	email := c.FormValue("email")
	company := c.FormValue("company")
	phone := c.FormValue("phone")

	fmt.Println(http.StatusOK, "id: "+id+", firstName: "+firstName+", lastName: "+lastName+", email: "+email+", company: "+company+", phone: "+phone)

	return c.JSON(http.StatusOK, []User{})
}

// swagger:route DELETE /v1/user/{id} users deleteUser
// Delete a user.
// responses:
//   200: deleteUserResponse

// deleteUser handles the DELETE request for deleting a user.
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "Deleted user in database with id="+id)
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
