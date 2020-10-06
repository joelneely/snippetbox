package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.

func home(w http.ResponseWriter, r *http.Request) {
	// Check whether the current request URL path exactly matches "/". If not,
	// use the http.NotFound() function to send a 404 response to the client.
	// Return from the handler to fail fast (without including the message).
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the id parameter's value from the query string,
	// converting it to an integer for security purposes via the
	// strconv.Atoi() function. On invalid data (or negative values),
	// return a 404 Not Found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Temporarily use fmt.Fprintf() to construct a response that
	// contains the snippet number (from the id parameter).
	// This replaces the original placeholder below:
	// w.Write([]byte("Display a specific snippet..."))
	fmt.Fprintf(w, "Display a specific snippet with id %d", id)
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Check whether the request used POST.
	if r.Method != http.MethodPost {
		// If not, add an 'Allow:POST' header to the response header map,
		// reject the request with "Method Not Allowed", then fail fast.
		w.Header().Set("Allow", http.MethodPost)
		// Instead of using separate calls...
		//   w.WriteHeader(405)
		//   w.Write([]byte("Method Not Allowed"))
		// ...the Error function does both at once.
		//
		// See also the book's description of setting application/json,
		// for direct access to the header map for exceptional names,
		// and for suppression of system-generated headers.
		//
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux;
	// register the handler functions for the corresponsing URL patterns.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we justcreated. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
