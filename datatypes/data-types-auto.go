package main

import (
	"fmt"
)

func main() {
	//	var cost int =1000
	cost := 1000000000000000000
	isCheck := false
	cmp := (5 + 12i)

	fmt.Printf("Type: %T  Value: %v  \n", cost, cost)
	fmt.Printf("Type: %T  Value: %v  \n", isCheck, isCheck)
	fmt.Printf("Type: %T  Value: %v  \n", cmp, cmp)
}
