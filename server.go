package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
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

type UserDB struct {
	ID        int    `db:"id"`
	FirstName string `db:"firstName"`
	LastName  string `db:"lastName"`
	Email     string `db:"email"`
	Company   string `db:"company"`
	Phone     string `db:"phone"`
}

// func connection() {

// 	// //connect to a PostgreSQL database
// 	// // Replace the connection details (user, dbname, password, host) with your own
// 	// db, err := sqlx.Connect("postgres", "user=postgres dbname=users sslmode=disable password=password host=localhost")
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// defer db.Close()

// 	// // Test the connection to the database
// 	// if err := db.Ping(); err != nil {
// 	// 	log.Fatal(err)
// 	// } else {
// 	// 	log.Println("Successfully connected to PostgreSQL database")
// 	// }

// 	// place := User{} // Initialize a User struct to hold retrieved data

// 	// // Execute a SQL query to select "username" and "email" columns from the "users" table
// 	// rows, _ := db.Queryx("SELECT * FROM users")

// 	// for rows.Next() {
// 	// 	err := rows.StructScan(&place) // Scan the current row into the "place" variable
// 	// 	if err != nil {
// 	// 		log.Fatalln(err)
// 	// 	}
// 	// 	log.Printf("%#v\n", place) // Log the content of the "place" struct for each row
// 	// }
// }

// swagger:route GET /v1/user users getUsers
// Get a list of users.
// responses:
//   200: usersResponse

// getUsers handles the GET request for retrieving a list of users.
func getUsers(c echo.Context) error {

	users := []User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com", Company: "ABC Inc.", Phone: "1234567890"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Company: "XYZ Corp.", Phone: "9876543210"},
		// Add more users as needed
	}

	// users := sq.Select("*").From("users").Join("emails USING (email_id)")
	// row := sqlx.DB.QueryRow("SELECT * FROM users")

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
	// row := sqlx.DB.QueryRow("INSERT INTO users (first_name, last_name, email, company, phone) VALUES ($1, $2, $3, $4, $5)", firstName, lastName, email, company, phone)

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
	// row := sqlx.DB.QueryRow("UPDATE users SET first_name=$1, last_name=$2, email=$3, company=$4, phone=$5 WHERE id=$6", firstName, lastName, email, company, phone, id)

	return c.JSON(http.StatusOK, []User{})
}

// swagger:route DELETE /v1/user/{id} users deleteUser
// Delete a user.
// responses:
//   200: deleteUserResponse

// deleteUser handles the DELETE request for deleting a user.
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	// row := sqlx.DB.QueryRow("DELETE FROM users WHERE id=$1", id)

	return c.JSON(http.StatusOK, "Deleted user in database with id="+id)
}

func main() {
	app := echo.New() // Create a new echo instance

	app.Use(middleware.CORS()) // Enable CORS

	port := ":1323"  // Port number
	version := "/v1" // API version

	//connect to the PostgreSQL database
	db, err := sqlx.Connect("postgres", "user=postgres dbname=users sslmode=disable password=password host=localhost")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("PostgreSQL connected successfully")
	}

	// routes
	app.POST(version+"/user", createUser)
	app.GET(version+"/user", getUsers)
	app.PUT(version+"/user/:id", updateUser)
	app.DELETE(version+"/user/:id", deleteUser)

	// start server
	app.Logger.Fatal(app.Start(port))
}
