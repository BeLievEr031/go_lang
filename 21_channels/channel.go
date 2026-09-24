package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	// fmt.Println("Processing the Num. ", <-numChan)
	for v := range numChan {
		fmt.Println("Processing the Num. ", v)
		time.Sleep(time.Second)
	}
}

func sendData(result chan int, num1 int, num2 int) {
	ans := num1 + num2
	result <- ans
}

func main() {
	// msgChannel := make(chan string)
	// msgChannel <- "sandeep"
	// msg := <-msgChannel
	// fmt.Println(msg)

	numChan := make(chan int)
	go processNum(numChan)
	for i := 0; i < 2; i++ {
		numChan <- rand.Intn(100)
	}

	result := make(chan int)
	go sendData(result, 5, 5)
	res := <-result

	fmt.Println("Result from channel: ", res)
}
