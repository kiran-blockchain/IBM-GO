package main

import (
	"fmt"
	"time"
)

func main() {
	
	go f("first")
	fmt.Println("I am main")
	
	time.Sleep(time.Second)
	fmt.Println("Done")
}

func f(inputStr string) {
	for i := 0; i < 3; i++ {
		fmt.Println(inputStr, ":", i)
	}
}
