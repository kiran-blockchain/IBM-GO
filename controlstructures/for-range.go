package main

import "fmt"

func main() {
	// Go lan integer array declaration
	nums := []int{2, 3, 4, 5}
	sum := 0
	//for-range is similar to for-each in javascript and in java
	for _, value := range nums {
		fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)
}
