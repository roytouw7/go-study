package Monad

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// TestMonad_Happy tests using the monad with function returning no error
func (test *TestSuite) TestMonad_Happy() {
	identity := func(i int) (int, error) {
		return i, nil
	}

	addOne := func(i int) (int, error) {
		return i + 1, nil
	}

	toString := func(i int) (string, error) {
		return fmt.Sprintf("%d", i), nil
	}

	a := Bind(Wrap(1), identity)
	b := Bind(a, addOne)
	c := Bind(b, toString)

	value, err := Unwrap(c)
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), "2", value)
}

// TestMonad_Happy tests using the monad with function returning error
func (test *TestSuite) TestMonad_Error() {
	identity := func(i int) (int, error) {
		return i, nil
	}

	addOneError := func(i int) (int, error) {
		return 0, fmt.Errorf("failed adding 1 to %d", i)
	}

	toString := func(i int) (string, error) {
		return fmt.Sprintf("%d", i), nil
	}

	a := Bind(Wrap(1), identity)
	b := Bind(a, addOneError)
	c := Bind(b, toString)

	value, err := Unwrap(c)
	assert.Error(test.T(), err)
	assert.Equal(test.T(), "", value) // default value for string
}

// TestMonad_BivariateFunction tests using the monad using a non-unary function (int, int) => (int, error)
func (test *TestSuite) TestMonad_BivariateFunction() {
	identity := func(i int) (int, error) {
		return i, nil
	}

	// currying to match throwable type definition
	sum := func(i int) func(j int) (int, error) {
		return func(j int) (int, error) {
			return i + j, nil
		}
	}

	a := Bind(Wrap(1), identity)
	b := Bind(a, sum(2))

	value, err := Unwrap(b)
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 3, value)
}

// TestMonad_ComplexType tests using the monad using a non-primary complex type
func (test *TestSuite) TestMonad_ComplexType() {
	type person struct {
		name string
		age  int
	}

	identity := func(i person) (person, error) {
		return i, nil
	}

	getPersonName := func(p person) (string, error) {
		return p.name, nil
	}

	p := person{
		name: "Roy",
		age:  30,
	}

	a := Bind(Wrap(p), identity)
	b := Bind(a, getPersonName)

	value, err := Unwrap(b)
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), "Roy", value)
}
