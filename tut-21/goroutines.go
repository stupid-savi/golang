package main

// Goroutines enables multi-threading

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1) // adding waitgroup
		go func(i int, w *sync.WaitGroup) {
			defer wg.Done() // removing waitgroup defer executes when func execution done like a clean up function
			fmt.Println("doing Task ", i)
		}(i, &wg)
	}

	wg.Wait() // wait till all threads are executed or waitgroup removed
}
