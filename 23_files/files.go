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

func readingFolder() {
	file, err := os.Open("../")

	if err != nil {
		panic(err)
	}

	folderInfo, err := file.ReadDir(-1)

	if err != nil {
		panic(err)
	}

	for _, fi := range folderInfo {
		fmt.Println(fi.Name())
	}
}

func constructFileFolderView() {
	file, err := os.Open("../")

	if err != nil {
		panic(err)
	}

	folderInfo, err := file.ReadDir(-1)

	if err != nil {
		panic(err)
	}

	for _, fi := range folderInfo {
		fmt.Println("|--", fi.Name())
		if fi.IsDir() && fi.Name() != ".git" {
			dirPath := "../" + fi.Name()
			constructView(dirPath, 4)
		}
	}

	defer file.Close()
}

func constructView(path string, space int) {
	folderInfo, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	fileInfo, err := folderInfo.ReadDir(-1)

	if err != nil {
		panic(err)
	}

	prefixSpace := ""

	for i := 1; i <= space; i++ {
		prefixSpace += " "
	}

	for _, fi := range fileInfo {
		fmt.Println(prefixSpace+"|--", fi.Name())
		if fi.IsDir() {
			dirPath := path + "/" + fi.Name()
			constructView(dirPath, space+4)
		}

	}

	defer folderInfo.Close()
}

func creatingFileWithWriteFN() {
	file, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}

	buff := []byte("hello go lang")

	n, err := file.Write(buff)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	defer file.Close()
}

func creatingFileWithWriteStringFN() {
	file, err := os.Create("example3.txt")

	if err != nil {
		panic(err)
	}

	file.WriteString("I am learning")
	file.WriteString(" go language.")

	defer file.Close()
}

func readingAndDumpingToAnotherFile() {
	// Open existing File data
	file, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()

	if err != nil {
		panic(err)
	}

	readBuff := make([]byte, fileInfo.Size())

	// Read existing File Data
	n, err := file.Read(readBuff)
	fmt.Println("File Size is: ", n)

	if err != nil {
		panic(err)
	}

	for i := range readBuff {
		fmt.Println(string(readBuff[i]))
	}

	// Create new file
	newFile, err := os.Create("new_file.txt")

	if err != nil {
		panic(err)
	}

	// Dump old file data
	newFile.Write(readBuff)

	defer func() {
		file.Close()
		newFile.Close()
	}()
}

func main() {

	// readingFileUsingBuffer()
	// readingFolder()
	// constructFileFolderView()
	// creatingFileWithWriteFN()
	// creatingFileWithWriteStringFN()
	// readingAndDumpingToAnotherFile()

	// Deleting Files

	// err := os.Remove("example2.txt")
	err := os.Remove("demo_1")
	// We can only delete the Empty folder.

	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted successfully.")
}
