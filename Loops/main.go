package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
}

func print(pi *int) { fmt.Println(*pi) }
