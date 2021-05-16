package models

// User is our user structure for authentication
type User struct {
	ID       int
	Login    string
	Password string
}
