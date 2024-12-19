package main

import (
	"fmt"
	"net/http"

	"github.com/cpwu/golang/pkg/handlers"
)

// Concepts
// Layouts

// Global port number for web server
const portNumber string = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting the application on port: %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)

}
