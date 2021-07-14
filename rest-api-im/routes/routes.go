package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	//Do Something
	ctx := r.Context()
   
	fmt.Fprintf(os.Stdout, "I am processing the ")

	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("Reuest processed"))
	case <-ctx.Done():
		fmt.Fprint(os.Stderr, "Cancelled")

	}

}
