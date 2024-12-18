package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Concepts
// How to serve html pages

// Global port number for web server
const portNumber string = ":8080"

// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl")
}

// Renders Templates
func renderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
