package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

// Mutex is a struct type and we create a zero valued variable m of type Mutex in line no. 15.
// In the above program we have changed the increment function so that the code which increments x x = x + 1 is between m.Lock() and m.Unlock().
// Now this code is void of any race conditions since only one Goroutine is allowed to execute this piece of code at any point of time.

// Now if this program is run, it will output

// final value of x 1000
