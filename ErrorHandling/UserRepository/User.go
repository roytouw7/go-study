package UserRepository

import (
	"errors"
	"fmt"
	"strconv"

	"Study/ErrorHandling/DataBase"
)

type InvalidUserDataError struct {
	invalidField string
	invalidValue string
}

func (e InvalidUserDataError) Error() string {
	return fmt.Sprintf("SetUserAge: could not set value %s: for field:%s", e.invalidValue, e.invalidField)
}

type DatabaseError struct {
	databaseError error
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("SetUserAge: database error: %w", e.databaseError)
}

func SetUserAge(userId string, age int) error {
	var (
		user *DataBase.User
		err  error
	)

	if err = checkAge(age); err != nil {
		return err
	}

	user, err = DataBase.GetUser(userId)
	if err != nil {
		var databaseConnError DatabaseError
		var userNotFoundError DataBase.UserNotFoundError

		switch {
		case errors.As(err, &databaseConnError):
			return DatabaseError{databaseError: err}
		case errors.As(err, &userNotFoundError):
			return InvalidUserDataError{
				invalidField: userId,
				invalidValue: "userId",
			}
		default:
			return fmt.Errorf("SetUserAge: unknown database error: %w", err)
		}
	}

	_, err = DataBase.SetUserAge(user, age)
	if err != nil {
		return fmt.Errorf("SetUserAge: failed setting user age: %w", err)
	}

	return nil
}

func checkAge(age int) error {
	if age > 100 || age < 0 {
		err := InvalidUserDataError{
			invalidValue: strconv.Itoa(age),
			invalidField: "age",
		}
		return err
	}

	return nil
}
