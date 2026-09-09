package main

func main() {
	// i := 10
	// switch i {
	// // Here we are using multiple condition.
	// case 0, 10:
	// 	println("Number found.")
	// default:
	// 	println("Number not found.")
	// }

	// switch time.Now().Weekday() {
	// case time.Sunday, time.Saturday:
	// 	println("It's a weekend mazza karo.")
	// default:
	// 	println("It's a working day.")
	// }

	whatIsMyType := func(i interface{}) {
		switch i.(type) {
		case int:
			println("I am integer.")
		case string:
			println("I am string.")
		default:
			println("Other.", i)
		}
	}

	whatIsMyType(10)
	whatIsMyType("Sandeep rajak")
	whatIsMyType(3.15)
}
