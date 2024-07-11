package models

// User represents a user in the system.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Company   string `json:"company"`
	Phone     string `json:"phone"`
}
