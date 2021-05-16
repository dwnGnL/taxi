package postgres

import (
	"fmt"
	"taxi/internal/auth/models"
)

func ExampleUserRepository() {
	url := "postgres://postgres:sql@localhost:5432/Travel"
	db := NewDatabase(url)

	userRepo := NewUserRepository(db)

	user := &models.User{
		Login:    "test",
		Password: "qwerty",
	}

	id, err := userRepo.SingIn(user)

	if err != nil {
		fmt.Println("no id")
		fmt.Println(err)
	} else {
		fmt.Println("id =", id)
		fmt.Println("could find the user")
	}

	// Output:
	// id = 1
	// could find the user
}
