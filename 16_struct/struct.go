// Structs are custome Data Structures available in Go.
// In Go we don't have Classes
package main

import (
	"fmt"
	"time"
)

type Student struct {
	name          string
	age           int
	percentage    float32
	isPass        bool
	admissionDate time.Time
}

func newStudent(name string, age int, percentage float32, isPass bool, admissionDate time.Time) *Student {
	newStudent := Student{
		name:          name,
		age:           age,
		percentage:    percentage,
		isPass:        isPass,
		admissionDate: admissionDate,
	}

	return &newStudent
}

// type Order struct {
// 	productName string
// 	price       float32
// 	status      string
// 	delivered   bool
// }

// func (o *Order) changeStatus(status string) {
// 	o.status = status
// }

// func changeDelivery(o *Order, delivered bool) {
// 	o.delivered = delivered
// }

// Learning struct embedding
// Basically passing a struct into another struct

type customerInfo struct {
	name    string
	pincode int
}

type Order struct {
	product string
	price   float32
	cutomer customerInfo
}

func main() {
	// order1 := Order{
	// 	productName: "Asus Tuf A15",
	// 	price:       84000,
	// 	status:      "PENDING",
	// 	delivered:   false,
	// }

	// fmt.Println(order1)
	// order1.changeStatus("PROCESSING")
	// fmt.Println(order1)

	// changeDelivery(&order1, true)
	// fmt.Println(order1)

	// student1 := newStudent("sandy", 24, 80.80, true, time.Now())

	// fmt.Println(student1.name)

	customer1 := customerInfo{
		name:    "sandeep",
		pincode: 41501,
	}

	order := Order{
		product: "Asus tuf a15",
		price:   84000.00,
		cutomer: customer1,
	}

	fmt.Println(order)
	fmt.Println("### Printing Customer name ###")
	fmt.Println(order.cutomer.name)
}
