package parsers

import "fmt"

type NegativeInt interface {
	Value() int
	isNegative()
}

type negativeInt int

func (n negativeInt) Value() int {
	return int(n)
}

func (n negativeInt) isNegative() {}

func ParseNegativeInt(i int) (NegativeInt, error) {
	if i > 0 {
		return nil, fmt.Errorf("can't parse %d as negative int", i)
	}

	return negativeInt(i), nil
}
