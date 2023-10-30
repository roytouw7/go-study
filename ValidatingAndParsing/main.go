package main

import (
	"fmt"

	"Study/ValidatingAndParsing/parsers"
)

type CustomInt int

// following the Parse Don't Validate design principle, I want at compile time to be sure the discount to be negative

// addDiscountAmount adds a discount amount to a sum
func addDiscountAmount(sum int, discount parsers.NegativeInt) int {
	return sum + discount.Value()
}

func identity(i CustomInt) CustomInt {
	return i
}

func main() {
	discount, err := parsers.ParseNegativeInt(-2)
	if err != nil {
		panic(err)
	}

	result := addDiscountAmount(7, discount)

	fmt.Println(result)

	identity(1) // I can pass 1 because it is implicitly casted to CustomInt
}
