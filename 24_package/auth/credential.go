package auth

import "fmt"

func LoginWitUsername(username string, password string) {
	fmt.Println("Username is: ", username)
	fmt.Println("Password is: ", password)
}

func authRule() {
	fmt.Println("Auth Rule")
}
