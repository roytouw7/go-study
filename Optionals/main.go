package main

import (
	"fmt"
	"strings"

	"Study/Optionals/Maybe"
)

func main() {
	Print("Hello World", nil)
	Print("Hello World", &PrintOptions{uppercase: true})
	Print2("Hello World", Maybe.Some(PrintOptions{uppercase: true}))
	var mOption Maybe.Maybe[PrintOptions] // investigate if this is preventable at build time
	Print2("Hello World", mOption)
	Print2("Hello World", Maybe.None[PrintOptions]())
}

type PrintOptions struct {
	uppercase bool
}

func Print(msg string, opt *PrintOptions) {
	if opt != nil {
		switch {
		case opt.uppercase:
			fmt.Println(strings.ToUpper(msg))
		}
	} else {
		fmt.Println(msg)
	}
}

func Print2(msg string, opt Maybe.Maybe[PrintOptions]) {
	options := opt.Value()
	switch {
	case options.uppercase:
		fmt.Println(strings.ToUpper(msg))
	default:
		fmt.Println(msg)
	}
}
