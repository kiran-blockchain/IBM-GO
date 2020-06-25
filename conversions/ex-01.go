package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"
	numberConverted, _ := strconv.Atoi(str)
	floatConverted, _ := strconv.ParseFloat(str, 8)
	fmt.Printf("%T \n", numberConverted)
	fmt.Printf("%T \n", floatConverted)
}
