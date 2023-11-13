package main

import (
	"errors"
	"fmt"
	"io/fs"
)

type OtherError struct {
	value        int
	Value2       string
	privateValue bool
}

func (e OtherError) Error() string {
	return "this is a silly error"
}

type TemplateParseError struct {
	value        int
	Value2       string
	privateValue bool
}

func (e TemplateParseError) Error() string {
	return fmt.Sprintf("my error %d", e.value)
}

func main() {
	var templateParseError TemplateParseError

	err := throws()
	if errors.Is(err, OtherError{7, "test", true}) {
		fmt.Println("expected")
	} else {
		fmt.Println("unexpected")
	}

	if errors.As(err, &templateParseError) {
		fmt.Println("expected 2")
	} else {
		fmt.Println("unexpected 2")
	}

	if errors.Is(errors.Unwrap(err), TemplateParseError{}) {
		fmt.Println("expected 3")
	} else {
		fmt.Println("unexpected 3")
	}

	errors.Is(err, fs.ErrNotExist)
}

func throws() error {
	err := causes()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func causes() error {
	return TemplateParseError{7, "test", true}
}
