package main

import "fmt"

func main() {
	// var arr [4]int

	// arr[0] = 10
	// arr[3] = 10
	// fmt.Println(arr)

	// arr2 := [5]int{10, 20, 30, 40, 50}

	// for i := 0; i < len(arr2); i++ {
	// 	println(arr2[i])
	// }

	arr3 := [4][3]int{{1, 2}, {3, 4}}
	row := len(arr3)
	column := len(arr3[0])

	// for i:=0;i<len()

	fmt.Println("Row Array len:: ", row)
	fmt.Println("Column Array len:: ", column)

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {

			if j != column-1 {
				fmt.Print(arr3[i][j], "-")
			} else {
				fmt.Print(arr3[i][j])

			}
		}
		fmt.Println()
	}

}
