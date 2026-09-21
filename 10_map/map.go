package main

import "fmt"

func main() {

	m := make(map[string]interface{})
	m["name"] = "Sandeep rajak"
	m["age"] = 45
	m["color"] = "red"
	m["isPass"] = true

	fmt.Println(m["name"])
	fmt.Println(m["isPass"])
	fmt.Println(m["color"])
	fmt.Println(m["age"])

	fmt.Println("Length of the MAP:: ")
	fmt.Println(len(m))

}
