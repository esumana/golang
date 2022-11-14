// Registering a Request Handler
// First, create a Handler which receives all incomming HTTP connections from browsers, HTTP clients or API requests. A handler in Go is a function with this signature:
//		func (w http.ResponseWriter, r *http.Request)
// The function receives two parameters:
// An http.ResponseWriter which is where you write your text/html response to.
// An http.Request which contains all information about this HTTP request including things like the URL
// or header fields.
// Registering a request handler to the default HTTP Server is as simple as this:

//		http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
//		    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
//		})

// The request handler alone can not accept any HTTP connections from the outside.
// The following code will start Go’s default HTTP server and listen for connections on port 80.

//		http.ListenAndServe(":80", nil)

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8000", nil)
}
