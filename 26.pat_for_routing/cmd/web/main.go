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
// Using Pat for Router
// Global port number for web server
const portNumber string = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	render.NewTemplates(&app)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting the application on port: %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
