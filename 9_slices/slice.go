package main

import (
	"fmt"
	"slices"
)

func main() {
	// var slice []int
	// fmt.Println(slice)
	// fmt.Println(slice == nil)

	var slice = make([]int, 2)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	//  CAP() Maximum element capacity

	var slice2 = make([]int, 1, 6)
	fmt.Println("Before Append function Data and Capacity.")
	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	slice2 = append(slice2, 1)
	slice2 = append(slice2, 2)
	slice2 = append(slice2, 3)
	slice2 = append(slice2, 4)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 6)

	fmt.Println("After Append function Data and Capacity.")
	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	// Copy Slice

	oslice := make([]int, 2)
	copy(slice, oslice)
	oslice = append(oslice, 10)

	fmt.Println("Prininting copied slice")
	fmt.Println(slice)
	fmt.Println(oslice)

	// Slices package
	// This package has multiple fn to work with multiple slice
	fmt.Println("Using functions from SLICES Construct")
	slice4 := []int{1, 2, 3, 4}
	slice5 := []int{1, 2, 3, 4}

	fmt.Println(slices.Equal(slice4, slice5))
	fmt.Println(slices.AppendSeq(slice4, slices.Values(slice5)))

	fmt.Println(slices.Values(slice5))

}
