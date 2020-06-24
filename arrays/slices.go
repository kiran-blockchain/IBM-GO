package main

import "fmt"

func main() {
	/* slice Declaration*/
	var s []int
	//var a = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s = a[2:10]
	s = append(s, 1000, 200, 10000, 10001, 1000001, 100101010)
	fmt.Println(s)

}
