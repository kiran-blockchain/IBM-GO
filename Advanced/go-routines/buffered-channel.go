package main

import (
	"fmt"
)

func main() {

	c := make(chan string)
	c <- "kiran"
	c <- "Ravi"
	fmt.Println(<-c)
	//fmt.Println(<-c)
}
