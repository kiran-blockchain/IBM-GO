package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Cannot Caluclate square root to a negative number")
	}
	return math.Sqrt(value), nil
}

func main() {
	x, err := sqrt(49.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}
