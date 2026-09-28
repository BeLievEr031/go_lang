package main

import (
	"fmt"
	"sync"
)

// In go the nutex is used for handling the Race condtion
// Race Condition: When some resource is updated by many users

type Post struct {
	view int
	mu   sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.view += 1

}

func main() {

	var wg sync.WaitGroup

	post1 := Post{
		view: 0,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go post1.inc(&wg)
	}

	wg.Wait()
	fmt.Println(post1.view)
}
