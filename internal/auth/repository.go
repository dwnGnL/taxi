package auth

import "taxi/internal/auth/models"

// UserRepository ...
type UserRepository interface {
	SingIn(request *models.User) (id int, err error)
	SignUp(request *models.User) (err error)
}

type Database interface {
	Close() error
}
