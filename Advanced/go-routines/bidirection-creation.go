package main

import (
	"fmt"
	"time"
)

func sendAPing(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}
func printForAPing(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}

}

func main() {
	mychan := make(chan string)
	go sendAPing(mychan)
	go printForAPing(mychan)

	var input string
	fmt.Scanln(&input)
}
