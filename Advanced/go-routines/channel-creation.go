package main

import "fmt"

func sending(s chan<- string) {
	//publishing the data to the channel
	s <- "Sending Data"
}
func main() {
	//creating a channel
	mychan := make(chan string)

	//enabled the go routine with channel
	go sending(mychan)

	//we will received the data published from the routine
	fmt.Println(<-mychan)
}
