package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
	"integra.com/go/cmd/models"
	"integra.com/go/cmd/storage"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// swagger:route GET /v1/user users getUsers
// Get a list of users.
// responses:
//
//	200: usersResponse
func GetUsers(c echo.Context) error {
	db := storage.GetDB()

	query := sq.Select("*").From("users")
	sql, _, _ := query.ToSql()

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Company, &user.Phone)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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
func CreateUser(c echo.Context) error {
	db := storage.GetDB()

	// obviously need to vet these inputs
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	company := c.FormValue("company")
	phone := c.FormValue("phone")

	dbName := os.Getenv("DB_NAME")

	sql, args, err := psql.Insert("").
		Into(dbName).
		Columns("firstname", "lastname", "email", "company", "phone").
		Values(firstname, lastname, email, company, phone).ToSql()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = db.Exec(sql, args...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"code": "201"})
}

// swagger:route PUT /v1/user/{id} users updateUser
// Update an existing user.
// responses:
//   200: userResponse

// updateUser handles the PUT request for updating an existing user.
func UpdateUser(c echo.Context) error {
	db := storage.GetDB()

	// obviously need to vet these inputs
	id := c.Param("id")

	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	company := c.FormValue("company")
	phone := c.FormValue("phone")

	dbName := os.Getenv("DB_NAME")

	sql, args, err := psql.Update(dbName).
		Set("firstname", firstname).
		Set("lastname", lastname).
		Set("email", email).
		Set("company", company).
		Set("phone", phone).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = db.Exec(sql, args...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"code": "200"})
}

// swagger:route DELETE /v1/user/{id} users deleteUser
// Delete a user.
// responses:
//   200: deleteUserResponse

// deleteUser handles the DELETE request for deleting a user.
func DeleteUser(c echo.Context) error {
	db := storage.GetDB()

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sql, args, err := psql.Delete("users").Where(sq.Eq{"id": idInt}).ToSql()
	if err != nil {
		return err
	}

	_, err = db.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "200"})
}
