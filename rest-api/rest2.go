package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/gorilla/mux"
)

func createServer() {
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/books", booksPage)
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homePage)
	// router.HandleFunc("/book", booksPage)
	router.HandleFunc("/countries", countriesPage)
	router.HandleFunc("/", myHandler)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/profile", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/", pprof.Index)

	http.ListenAndServe(":4000", router)
}
func booksPage(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "<h1>Welcome to Go Lang Articles</h1>")
}
func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Welcome to GoLang Training</h1>")
}


func myHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("MyResponse"))
}

func main() {
	createServer()
}

func countriesPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, receiveData())

}
func receiveData() (countries string) {
	response, err := http.Get("https://restcountries.eu/rest/v2/all")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	countries = string(responseData)
	return
}

// func profiler(){

// }
