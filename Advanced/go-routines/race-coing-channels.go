package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	//channel will be set to true
	ch <- true
	x = x + 1
	// channel will be rest to false
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	//buffered channel with the capacity as one.
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
