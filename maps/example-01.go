package main

import "fmt"

func main() {
	var m = make(map[string]string)
	m["emp01"] = "Kiran"
	m["emp02"] = "Ravi"
	for key, value := range m {
		fmt.Println("key:", key, "Value:", value)
	}
	delete(m, "emp01")
	for key, value := range m {
		fmt.Println("key:", key, "Value:", value)
	}

	i, t := m["emp02"]
	fmt.Println("isExisting:", i, "Value-t:", t)
}
