package main

import (
	"fmt"
)

func main(){
	fmt.Println("I am good1")
	defer fmt.Println("Under Stand The invokin g defer1")
	defer fmt.Println("Under Stand The invokin g defer2")
	fmt.Println("I am good2")
	fmt.Println("I am good3")
}

