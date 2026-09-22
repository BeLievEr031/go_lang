package main

import "fmt"

func variadicFun(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func main() {

	ans := variadicFun(1, 2, 3, 4, 5, 6)
	slc := []int{1, 2, 3}
	res := variadicFun(slc...)
	fmt.Println(ans)
	fmt.Println(res)
}
