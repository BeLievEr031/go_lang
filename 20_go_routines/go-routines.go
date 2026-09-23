package main

import (
	"fmt"
	"sync"
)

// func task(id int) {
// 	fmt.Println(id)
// }

func wgTask[T interface{}](id T, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(id)
}

// func deferFn() {
// 	defer func() {
// 		fmt.Println("I will run after this function completion")
// 	}()

// 	fmt.Println("Starting fn.")
// 	fmt.Println("Processing fn.")
// 	fmt.Println("Ending fn.")
// }

func main() {
	// for i := 1; i <= 10; i++ {
	// 	go task(i)
	// }

	// time.Sleep(time.Second * 2)

	// deferFn()

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go wgTask(i, &wg)
	}

	wg.Wait()
}
