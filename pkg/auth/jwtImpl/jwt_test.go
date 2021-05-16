package jwtImpl

import "fmt"

func ExampleJWT() {
	signingKey := "test"
	var expireDuration int64 = 5

	JWT, err := NewJWT(signingKey, expireDuration)

	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("created")
	}

	token, err := JWT.NewJWT(5)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("token was created")
	}

	id, err := JWT.ParseToken(token)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("no error occurred")
	}
	fmt.Println("id =", id)

	// Output:
	// created
	// token was created
	// no error occurred
	// id = 5

}
