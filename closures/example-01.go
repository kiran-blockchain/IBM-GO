	package main

	import (
		"fmt"
	)
	//outer function
	func intSequence() func() int {
		a := 0
		//inner function
		return func() int {
			a++
			return a
		}
	}

func main() {
	getnext := intSequence()
	
	second := intSequence()
	
	fmt.Println(getnext())
	fmt.Println(getnext())
	fmt.Println(getnext())

	fmt.Println(second())
	fmt.Println(second())
	fmt.Println(second())

}
