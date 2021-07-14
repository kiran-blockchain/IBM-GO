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
	//This will execute last
	defer deferMain()
	//This one will execute first
	defer deferOne()
	//fmt.Println("I am main")
	fmt.Println("I am before the panic")
	panic("Help I am a fatal error!!!!!")
}
