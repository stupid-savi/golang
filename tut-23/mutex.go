// Mutexs are used to avoid race condition during multi-threading
// They can lock the particular logic or line

package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) incViews(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock() // Always try to lock minimum lines not entire function do threadings need to do less work when unlocks
	p.views += 1

}

func main() {

	var wg sync.WaitGroup

	myPost := post{
		views: 0,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.incViews(&wg) // This leads to race condition
	}

	wg.Wait()
	fmt.Println("Views", myPost.views)

}
