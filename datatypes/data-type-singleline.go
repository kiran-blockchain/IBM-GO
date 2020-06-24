package main

import (
	"fmt"
)

func main() {
	//	var cost int =1000
	const name = "KIRAN"
	cost, price := 1000, 2000
	isCheck := false
	cmp := (5 + 12i)
	//name = "RAVI"

	fmt.Printf("Type: %T  Value: %v  \n", cost, cost)
	fmt.Printf("Type: %T  Value: %v  \n", isCheck, isCheck)
	fmt.Printf("Type: %T  Value: %v  \n", cmp, cmp)
	fmt.Printf("Type: %T  Value: %v  \n", price, price)
	fmt.Printf("Type: %T  Value: %v  \n", name, name)
}
