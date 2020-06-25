package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}


func (emp Employee) getDetails(){
	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
}

func main() {
	e := Employee{Name: "Kiran", Age: 30}
	e.getDetails()
}
