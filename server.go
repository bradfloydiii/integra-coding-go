package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

// User represents a user in the system.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Company   string `json:"company"`
	Phone     string `json:"phone"`
}

var db *sql.DB

func main() {
	app := echo.New()          // Create a new echo instance
	app.Use(middleware.CORS()) // Enable CORS

	port := ":1323"  // Port number
	version := "/v1" // API version

	// connect to the postgres database
	var err error
	db, err = sql.Open("postgres", "user=postgres password=password dbname=users port=5432 sslmode=disable") // would separate this into a config file
	if err != nil {
		log.Fatal(err)
	}

	// test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("database connected successfully")
	}

	app.GET(version+"/user", getUsers)
	app.POST(version+"/user", createUser)
	app.PUT(version+"/user/:id", updateUser)
	app.DELETE(version+"/user/:id", deleteUser)

	app.Logger.Fatal(app.Start(port))
}

// swagger:route GET /v1/user users getUsers
// Get a list of users.
// responses:
//
//	200: usersResponse
func getUsers(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Company, &user.Phone)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, users)
}

// swagger:route POST /v1/user users createUser
// Create a new user.
// responses:
//
//	200: userResponse
//
// createUser handles the POST request for creating a new user.
func createUser(c echo.Context) error {

	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	company := c.FormValue("company")
	phone := c.FormValue("phone")

	fmt.Println("firstname: " + firstname + ", lastname: " + lastname + ", email: " + email + ", company: " + company + ", phone: " + phone)

	stmt, err := db.Prepare("INSERT INTO users (firstname, lastname, email, company, phone) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(firstname, lastname, email, company, phone); err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "User NOT created successfully"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
}

// swagger:route PUT /v1/user/{id} users updateUser
// Update an existing user.
// responses:
//   200: userResponse

// updateUser handles the PUT request for updating an existing user.
func updateUser(c echo.Context) error {
	id := c.Param("id")
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	company := c.FormValue("company")
	phone := c.FormValue("phone")

	fmt.Println("id: " + id + ", firstname: " + firstname + ", lastname: " + lastname + ", email: " + email + ", company: " + company + ", phone: " + phone)

	stmt, err := db.Prepare("UPDATE users SET firstname=$1, lastname=$2, email=$3, company=$4, phone=$5 WHERE id=$6")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(firstname, lastname, email, company, phone, id); err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "User NOT updated successfully"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
}

// swagger:route DELETE /v1/user/{id} users deleteUser
// Delete a user.
// responses:
//   200: deleteUserResponse

// deleteUser handles the DELETE request for deleting a user.
func deleteUser(c echo.Context) error {
	id := c.Param("id")

	fmt.Println("id: " + id)

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec(id); err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "User NOT deleted successfully"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted user successfully"})
}
