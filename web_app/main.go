package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main application fucntion
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting server on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
