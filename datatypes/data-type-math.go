package main

import (
	"fmt"
	"math"
)

func main() {
	//	var cost int =1000
	var cost float64 = 900
	sqrt := math.Sqrt(cost)

	fmt.Printf("Type: %T  Value: %v  \n", cost, cost)
	fmt.Printf("Type: %T  Value: %v  \n", sqrt, sqrt)

}
