package main

import (
	"fmt"
)

//outer function
func varidaicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[1])
}

func calculation(s string, val ...int) {
	if s == "rect" {
		fmt.Println(val[0] * val[1])
	}
	if s == "square" {
		fmt.Println(val[0] * val[0])
	}
}

func main() {
	varidaicExample("Kiran", "Ravi")
	calculation("rect", 20, 30)
	calculation("square", 20)
}
