package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()

	if err != nil {
		panic(err)
	}

	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("File size: ", fileInfo.Size())
	fmt.Println("File Mode: ", fileInfo.Mode())
	fmt.Println("File IsDir: ", fileInfo.IsDir())
	fmt.Println("File Modified time: ", fileInfo.ModTime())
	// fmt.Println("File Modified time: ", fileInfo.ModTime().Add(time.Hour))
	// fmt.Println("File Modified time: ", fileInfo.Sys())

}
