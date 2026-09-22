package main

import "fmt"

func main() {
	fmt.Println("i am range file.")
	slc := make([]int, 2)
	slc[0] = 1
	slc[1] = 2

	slc = append(slc, 3)
	slc = append(slc, 4)
	slc = append(slc, 5)
	slc = append(slc, 6)
	// fmt.Println(slc)

	// for i := 0; i < len(slc); i++ {
	// 	fmt.Println(slc[i])
	// }
	sum := 0
	for _, elem := range slc {
		sum += elem
	}

	// var mp = map[string]int{"red": 1, "blue": 2}

	// fmt.Println(mp)

	// for key := range mp {
	// 	fmt.Println(key, ":", mp[key])
	// }

	// Starting Byte of rune
	// 255 -> 1 Byte
	for k, v := range "sandeep" {
		fmt.Println(k, string(v))
	}
}
