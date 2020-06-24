package main

import "fmt"

func main() {
	/* slice Declaration*/
	var s []int

	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = a[0:5]
	//s = append(s, 1000, 200, 10000, 10001, 1000001, 100101010)
	fmt.Println(cap(s))

	//Best practices
	var c = make([]int, 3)
	copy(c, a[2:5])
	fmt.Println(c)
	fmt.Println(cap(c))
	//delete an element from slice
	c = append(c[:1], c[2])
	fmt.Println(c)
}
