package main

import (
	"fmt"
	"taxi/internal/app"
)

func main() {

	fmt.Println("Starting the programm ...")

	confpath := "config/config.json"

	err := app.Run(confpath)
	if err != nil {
		fmt.Println(err)
	}
}
