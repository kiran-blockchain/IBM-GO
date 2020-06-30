package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Getenv("myenv")
	fmt.Printf("%s \n", env)
}
