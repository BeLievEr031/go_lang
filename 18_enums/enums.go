// In GO Lang we do not have any specific thing like ENUM
// But we can create enum using const
package main

import "fmt"

type Role int

const (
	Admin Role = iota + 1
	SuperAdmin
	User
)

type OrderStatus string

const (
	Pending    OrderStatus = "Pending"
	Processing OrderStatus = "Processing"
	Completed  OrderStatus = "Completed"
)

func seeRole(role Role) {
	fmt.Println("You have selected: ", role)
}

func checkStatus(status OrderStatus) {
	fmt.Println("You status is: ", status)
}

func main() {
	seeRole(Admin)
	seeRole(User)

	checkStatus(Pending)
	checkStatus(Completed)
}
