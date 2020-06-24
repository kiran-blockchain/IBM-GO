package main

import (
	"fmt"
)

//outer function
func main() {
	a := 10
	var p *int = &a
	//inner function
	fmt.Println("Address of a is ", &a)
	fmt.Println("Value of a is ", a)

	fmt.Printf("Type of p is %T \n", p)
	fmt.Println("Address of p is  \n", p)
}
