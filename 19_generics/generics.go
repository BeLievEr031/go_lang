package main

import "fmt"

func printIntSlc(slc []int) {
	for _, v := range slc {
		fmt.Println(v)
	}
}

func printStrSlc(slc []string) {
	for _, v := range slc {
		fmt.Println(v)
	}
}

func printSlc[T int | string](slc []T) {
	for _, v := range slc {
		fmt.Println(v)
	}
}

type Person struct {
	name string
	slc  []int
}

type PersonWithGenerics[T any] struct {
	name string
	slc  []T
}

func main() {

	// printIntSlc([]int{1, 2, 3, 4})
	// printStrSlc([]string{"sandy", "candy", "mandy"})
	printSlc([]int{1, 2, 3})
	printSlc([]string{"sandy", "sandeep"})
	// printSlc([]bool{true, false})
	person1 := Person{
		name: "sandeep",
		slc:  []int{1, 2, 3},
	}

	person2 := PersonWithGenerics[int]{
		name: "sandy",
		slc:  []int{1, 2, 3, 4},
	}

	person3 := PersonWithGenerics[string]{
		name: "sandy-031",
		slc:  []string{"One", "Two"},
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
}
