package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	handleRequests()
	fmt.Println("Welcome to tracing")
	var wg sync.WaitGroup
	wg.Add(1)
	go checkMemoryLeaks(wg)
	wg.Wait()
}
func handleRequests() {

	// myRoutes := mux.NewRouter()
	// myRoutes.HandleFunc("/", indexPage).Methods("GET")
	// myRoutes.HandleFunc("/getBooks", getBooks).Methods("POST")

	fmt.Println("Server running on port-3000")
	//	http.HandleFunc("/debug/pprof/profile", pprof.Profile)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
func checkMemoryLeaks(wg sync.WaitGroup) {

	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "Test functionality")
		if (i % 10000) == 0 {
			time.Sleep(500 * time.Second)
		}
	}

}
