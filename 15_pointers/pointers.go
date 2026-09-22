package main

import "fmt"

func changeValue(num *int) {
	*num = 45 // De-Referencing for changing the value
	fmt.Println("Num value in func: ", *num)
}

func main() {
	num := 1
	changeValue(&num)
	fmt.Println("Num value outside func: ", num)
}
