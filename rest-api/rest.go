package main

import (
	"fmt"
	"log"
	"net/http"
)

//render a index page.
func indexPage(res http.ResponseWriter, req *http.Request) {
	fmt.Println("I got the request")
	fmt.Fprintf(res, "Welcome to GoLang Rest Server")
	fmt.Println("I sent the response")
}

//Handle Requests
func handleRequests() {
	http.HandleFunc("/", indexPage)
	fmt.Println("Server running on port-3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func main() {
	handleRequests()
}
