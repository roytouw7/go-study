package Monad

import (
	"context"
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

	// manually currying to match unaryThrower type definition
	sum := func(i int) func(j int) (int, error) {
		return func(j int) (int, error) {
			return i + j, nil
		}
	}

	// normal implementation, easier to use
	sumNormal := func(i, j int) (int, error) {
		return i + j, nil
	}

	a := Bind(Wrap(1), identity)
	b := Bind(a, sum(2))
	c := Bind(b, Curry(2, sumNormal))

	value, err := Unwrap(c)
	assert.Nil(test.T(), err)
	assert.Equal(test.T(), 5, value)
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

type HotelService struct{}

type Hotel struct {
	uuid string
}

type Policies struct {
	policy string
}

type WebPages struct {
	page string
}

func (h *HotelService) GetHotel(ctx context.Context, uuid string) (Hotel, error) {
	return Hotel{uuid: uuid}, nil
}
func (h *HotelService) GetGuaranteePolicies(ctx context.Context, hotel Hotel) (Policies, error) {
	return Policies{policy: fmt.Sprintf("Policy for hotel %s", hotel.uuid)}, nil
}
func (h *HotelService) GetWebpagesForHotel(ctx context.Context, hotel Hotel) (WebPages, error) {
	return WebPages{page: fmt.Sprintf("Page for hotel %s", hotel.uuid)}, nil
}

var hotelService HotelService

// TestMonad_RealLifeExample_WithoutMonad tests the real-life example without a monad, as found in real code
func (test *TestSuite) TestMonad_RealLifeExample_WithoutMonad() {
	hotel, err := hotelService.GetHotel(nil, "1337")
	if err != nil {
		return // return error
	}

	policies, err := hotelService.GetGuaranteePolicies(nil, hotel)
	if err != nil {
		return // return error
	}

	webpages, err := hotelService.GetWebpagesForHotel(nil, hotel)
	if err != nil {
		return // return error
	}

	test.Equal(policies.policy, fmt.Sprintf("Policy for hotel %s", hotel.uuid))
	test.Equal(webpages.page, fmt.Sprintf("Page for hotel %s", hotel.uuid))
}

// TestMonad_RealLifeExample_WithMonad tests the real-life example using a monad
func (test *TestSuite) TestMonad_RealLifeExample_WithMonad() {
	hotel := Bind(Wrap("1337"), Curry(nil, hotelService.GetHotel))
	policies := Bind(hotel, Curry(nil, hotelService.GetGuaranteePolicies))
	webpages := Bind(hotel, Curry(nil, hotelService.GetWebpagesForHotel))

	if webpages.Err != nil { // todo not a fan of this, appears we only check webpages error now
		return // return error	// benefit might be bigger if the error handling is more than just a return err
	}

	test.Equal(policies.Value.policy, fmt.Sprintf("Policy for hotel %s", hotel.Value.uuid))
	test.Equal(webpages.Value.page, fmt.Sprintf("Page for hotel %s", hotel.Value.uuid))
}

// todo is there ever a case for nullary functions where a monad is nice? not really? you could just wrap with the error en join the monad from there? make unit test to proof this works
