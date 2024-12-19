package main

import (
	"fmt"
	"net/http"

	"github.com/cpwu/golang/pkg/handlers"
)

// Concepts
// Enable go modules + reorganize for go modules + templates from layout + passing data.

// Global port number for web server
const portNumber string = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
