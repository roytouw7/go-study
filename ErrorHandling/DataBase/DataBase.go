package DataBase

import "fmt"

type User struct {
	id   string
	Name string
	Age  int
}

type DatabaseConnError struct{}

func (e DatabaseConnError) Error() string {
	return "failed connecting to the database"
}

type UserNotFoundError struct {
	userId string
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("user with id %s not found in database", e.userId)
}

// TODO make methods to test pointer receiver handling

func GetUser(userId string) (*User, error) {
	if userId != "1" {
		return nil, UserNotFoundError{userId: userId}
	}

	return &User{
		id:   userId,
		Name: "Roy",
		Age:  30,
	}, nil
}

func SetUserAge(user *User, age int) (*User, error) {
	user.Age = age
	return user, nil
}
