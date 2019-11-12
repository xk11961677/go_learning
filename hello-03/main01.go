/**
并发
*/
package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	total int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		total.Lock()
		total.total++
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.total)
}
