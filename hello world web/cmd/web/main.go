package main

import (
	"fmt"
	"net/http"

	"github.com/tsawler/go-course/pkg/handlers"
)

var port = "8080"

// if the first letter of a function is lowercase
// the function is not accessible elsewhere from outside the package

// if its uppercase, it can be accessible from outside the package

/*
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			log.Println(err)
		}

		fmt.Println(fmt.Sprintf("Bytes writted: %d", n))
	})
*/

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	_ = http.ListenAndServe(":"+port, nil)

}
