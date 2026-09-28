package main

import (
	"fmt"
	"os"
)

func basicFileInfoFunc() {
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
}

func readingFileUsingBuffer() {
	file, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	// buff := make([]byte, 25) Bad Approach,
	// Note: If extra space was assigned that remain empty do not print, but memory still consumes

	fileInfo, err := file.Stat()

	buff := make([]byte, fileInfo.Size())
	ln, err := file.Read(buff)

	for i := 0; i < ln; i++ {
		fmt.Println("Buffer data: ", string(buff[i]))
	}

	fmt.Println(ln)
	defer file.Close()
}

func main() {

	readingFileUsingBuffer()

}
