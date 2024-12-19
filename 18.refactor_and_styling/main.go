package main

import (
	"fmt"
	"net/http"
)

// Concepts
// Refactor into distinct go files for specific purposes + bootstrap

// Global port number for web server
const portNumber string = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
