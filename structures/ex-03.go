package main

import (
	"fmt"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {
	r1 := rectangle{length: 100, breadth: 200, color: "yellow"}
	fmt.Println(r1)
	//Call by value
	r2 := r1
	r2.color = "Pink"
	fmt.Println("-----------")
	fmt.Println("r2:", r2)
	fmt.Println("r1", r1)

	//Cal by Ref
	r3 := &r1
	r3.color = "Green"
	fmt.Println("-----------")
	fmt.Println("r3:", r3)
	fmt.Println("r1", r1)

}
