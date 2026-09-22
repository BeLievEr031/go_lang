package main

import "fmt"

func add(a int, b int) int {
	sum := a + b
	return sum
}

func multiReturn() (string, int, bool) {
	return "sandeep", 24, true
}

func fruitInfo() (string, string, int, bool) {
	return "mango", "yellow", 120, true
}

func processAmt(cb func(disount int) int, paid bool) {
	if paid {
		res := cb(6)
		fmt.Println("Discounted Amount is: ", res)
	}
}

func discountFunc(dis int) int {
	return dis * 120
}

func returnFn() func(a int) int {
	return func(b int) int {
		return b
	}
}

func main() {
	fmt.Println(add(5, 6))

	name, age, isPass := multiReturn()
	fmt.Println(name, age, isPass)

	fruitName, _, price, _ := fruitInfo()
	fmt.Println(fruitName, price)

	processAmt(discountFunc, true)

	fn := returnFn()
	fn(5)
}
