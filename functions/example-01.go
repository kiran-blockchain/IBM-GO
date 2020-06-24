package main

import "fmt"

func main() {
	_, y := calculate(10, 20)
	//fmt.Println(x)
	fmt.Println(y)
}
func calculate(a int, b int) (x int, y int) {
	x = a + b
	y = a * b
	return
}
