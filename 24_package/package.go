package main

import (

	// "image/color"

	"github.com/BeLievEr031/go_lang/auth"
	"github.com/fatih/color"
)

func main() {

	auth.LoginWitUsername("sandy", "password123")
	session := auth.Session()

	// fmt.Println(session)
	// color.BlinkRapid = 1
	color.Yellow(session)
	// color.BgGreen(session)
}
