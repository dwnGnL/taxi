package postgres

import (
	"context"
	"errors"
	"fmt"
	"taxi/internal/auth/models"
)

var errUserExist = errors.New("user already exists")

type user struct {
	ID       int
	Login    string
	Password string
}

func toPostgresUser(u *models.User) *user {
	return &user{
		ID:       u.ID,
		Login:    u.Login,
		Password: u.Password,
	}
}

func toModel(u *user) *models.User {
	return &models.User{
		ID:       u.ID,
		Login:    u.Login,
		Password: u.Password,
	}
}

// UserRepository ...
type UserRepository struct {
	db *Database
}

// NewUserRepository ...
func NewUserRepository(db *Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// SingIn ...
func (user *UserRepository) SingIn(request *models.User) (id int, err error) {
	fmt.Println("(user *UserRepository) SingIn start")
	pgUser := toPostgresUser(request)

	err = user.db.conn.QueryRow(context.Background(), "select id from users where login=$1 and pass=$2", pgUser.Login, pgUser.Password).Scan(&id)
	if err != nil {
		fmt.Printf("QueryRow failed: %v\n", err)
	}
	fmt.Println("(user *UserRepository) SingIn end")
	return id, err
}

// SignUp ...
func (user *UserRepository) SignUp(request *models.User) (err error) {

	pgUser := toPostgresUser(request)
	var id uint64
	err = user.db.conn.QueryRow(context.Background(), "select id from users where login=$1", pgUser.Login).Scan(&id)
	if err != nil {
		if err == errNoRows {
			err = user.db.conn.QueryRow(context.Background(), "INSERT INTO USERS(login, pass) VALUES($1, $2) RETURNING id", request.Login, request.Password).Scan(&id)
			if err != nil {
				fmt.Printf("Unable to INSERT: %v\n", err)
				return err
			}
			return nil
		}
		fmt.Printf("QueryRow failed: %v\n", err)
	}
	return
}
