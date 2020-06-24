package main

import "fmt"

func main() {

	var i int
	fmt.Scanln(&i)

	switch i {
	case 10:
		fmt.Print("The value is 10")
	case 20:
		fmt.Print("The value is 20")
	default:
		fmt.Print(i)
	}
}
