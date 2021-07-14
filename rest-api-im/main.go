package main

import (
	"fmt"
	"net/http"
	"rest-api/routes"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Go Lang Http server")
	port := ":4000"
	r := mux.NewRouter()

	r.HandleFunc("/", routes.DefaultRoute).Methods("GET")
	//http.HandleFunc("/", deafultHandler)
	http.ListenAndServe(port, r)
}

func deafultHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Fprintf(res, "Hello Guys Welcome to POST request")
	} else {
		fmt.Fprintf(res, "Hello Guys Welcome to GET request")
	}

}
