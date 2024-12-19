package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cpwu/golang/pkg/config"
	"github.com/cpwu/golang/pkg/handlers"
	"github.com/cpwu/golang/pkg/render"
)

// Concepts
// Template Cache

// Global port number for web server
const portNumber string = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CeateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting the application on port: %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)

}
