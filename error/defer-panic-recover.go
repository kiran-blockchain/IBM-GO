package main

import (
	"fmt"
)

func deferOne() {
	fmt.Println("deferone Started")
	if r := recover(); r != nil {
		fmt.Println(r)
		fmt.Println("I have a panic situation")
		//defer deferOne()
	}
	fmt.Println("deferone done")

}

func deferMain() {
	fmt.Println("defer Main Started")
	if r := recover(); r != nil {
		fmt.Println(r)
		fmt.Println("I have a panic situation")
		//defer deferOne()
	}
}
func main() {

	defer deferMain()
	defer deferOne()
	//fmt.Println("I am main")

	panic("Help!!!!!")
}
