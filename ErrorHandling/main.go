package main

import (
	"errors"
	"fmt"

	"Study/ErrorHandling/UserRepository"
)

func main() {
	var (
		user = "2"
		age  = 77
	)
	err := UserRepository.SetUserAge(user, age)
	if err != nil {
		var invalidUserDataError UserRepository.InvalidUserDataError

		switch {
		case errors.As(err, &invalidUserDataError):
			fmt.Println("invalid user data, %w", err)
		default:
			panic("unknown error state")
		}
	} else {
		fmt.Printf("set user with ID: %s to age:%d", user, age)
	}
}
