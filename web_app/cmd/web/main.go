package main

import (
	"fmt"
	"net/http"

	"github.com/cpwu/golang/web_app/handlers"
)

const portNumber = ":8080"

// main is the main application fucntion
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting server on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
