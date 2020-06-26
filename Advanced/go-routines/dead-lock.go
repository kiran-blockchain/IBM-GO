package main

import (
	"fmt"
)

func main() {
	mesages := make(chan string)
	go func() {
		fmt.Println("I am trying to recevie messages")//printing
		<-mesages//blockcing
		fmt.Println("Receivie Messages")//it will never reacher
	}()
	//mesages <- "i am dead lock"
	fmt.Println("I am trying to recevie messages")//printing
	<-mesages
	fmt.Println("Receivie Messages")//it will never reacher
}
