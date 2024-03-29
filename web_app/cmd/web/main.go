package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cpwu/golang/web_app/pkg/render"

	"github.com/cpwu/golang/web_app/config"

	"github.com/cpwu/golang/web_app/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application fucntion
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache.")
	}

	app.TemplateCache = templateCache
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting server on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
