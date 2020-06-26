package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book Strcuture
type Book struct {
	Title   string `json:Title`
	Author  string `json:Author`
	Content string `json:Content`
}

//Books list
var Books []Book

//render a index page.
func indexPage(res http.ResponseWriter, req *http.Request) {
	fmt.Println("I got the request")
	fmt.Fprintf(res, "Welcome to GoLang Rest Server")
	fmt.Println("I sent the response")
}

func getBooks(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(Books)
}

//Handle Requests
func handleRequests() {
	
	myRoutes := mux.NewRouter()
	myRoutes.HandleFunc("/", indexPage).Methods("GET")
	myRoutes.HandleFunc("/getBooks", getBooks).Methods("POST")
	
	fmt.Println("Server running on port-3000")
	
	log.Fatal(http.ListenAndServe(":3001", nil))

}
func handleRoutes() {
	myRoutes := mux.NewRouter()
	myRoutes.HandleFunc("/", indexPage).Methods("GET")
	myRoutes.HandleFunc("/getBooks", getBooks).Methods("POST")
}
func main() {
	Books = []Book{
		Book{Author: "Kiran", Title: "Go Languag", Content: "This is awesome"},
		Book{Author: "Sampath",
			Title:   "Advanced Go Languag",
			Content: "This is awesome too"},
	}
	handleRequests()
}
