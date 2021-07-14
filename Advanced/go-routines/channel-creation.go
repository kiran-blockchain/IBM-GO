package main

import (
	"fmt"
	"time"
)

func sending(s chan<- string) {
	fmt.Println("Sendign Data1")
	//publishing the data to the channel
	s <- "Sending Data 1"
	
}

func sending2(s chan<- string) {
	//publishing the data to the channel
	fmt.Println("Sendign Data2")
	s <- "Sending Data 2"
}
func main() {
	//creating a channel
	mychan := make(chan string)

	//enabled the go routine with channel
	go sending(mychan)
	go sending2(mychan)

	//we will received the data published from the routine
	fmt.Println(<-mychan)
	fmt.Println(<-mychan)
	time.Sleep(time.Second * 5)
}
