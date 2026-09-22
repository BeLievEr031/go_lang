package main

import "fmt"

func counter() func() int {
	ct := 0

	return func() int {
		ct += 1
		return ct
	}
}

func main() {
	incFn := counter()

	fmt.Println(incFn())
	fmt.Println(incFn())
	fmt.Println(incFn())
	fmt.Println(incFn())
	fmt.Println(incFn())
}
