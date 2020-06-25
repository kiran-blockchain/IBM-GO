package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//message := "Welcome to Go Lang"
	filename := "./fileone.txt"
	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	inputData := string(filebuffer)
	fmt.Println(inputData)
}
