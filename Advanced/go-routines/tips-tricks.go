//count the number of go routines

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	go func() {
	}()
	go func() {
	}()
	go func() {
	}()
	fmt.Println(runtime.NumGoroutine())
}
