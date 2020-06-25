package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	jsonString := `
 {
	 "firstname":"kiran",
	 "lastname":"pvs",
	 "age":30
 }`
	emp1 := new(Employee)
	json.Unmarshal([]byte(jsonString), emp1)
	fmt.Println(emp1.FirstName)
	fmt.Println(emp1.LastName)
	fmt.Println(emp1.Age)

	emp2 := Employee{FirstName: "JOHN", LastName: "Paul", Age: 50}
	jsonFormat, err := json.Marshal(emp2)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(jsonFormat)
		fmt.Printf("%s \n", jsonFormat)
	}

}
