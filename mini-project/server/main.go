package main

import (
	"fmt"
	"log"
	"net/http"

	"./router"
)

func main() {
	r := router.Router()
	fmt.Println("Server running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
