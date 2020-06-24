package main

import "fmt"

func main() {
	/* Array Declaration*/
	var employees = [3][3]string{{"1", "2", "3"},
		{"Kiran", "Rupsh", "Aditi"},
		{"JOHN", "krish", "lll"}}
	fmt.Println(employees[1][0])
	for index, value := range employees {
		fmt.Println(index)
		fmt.Println(value)

		for index2, value2 := range value {
			fmt.Println(index2)
			fmt.Println(value2)
		}
	}
}
