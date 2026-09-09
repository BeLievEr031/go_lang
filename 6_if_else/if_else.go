package main

func main() {
	println("if else tutorial")

	age := 175

	if age < 18 {
		println("You are teenager.")
	} else if age <= 60 {
		println("Adult")
	} else {
		println("Senior citizen")
	}
}
