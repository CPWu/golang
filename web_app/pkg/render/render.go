package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders tempates using html/template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if we already have the template
	_, inMap := templateCache[t]
	if !inMap {
		// Need to create the template
		log.Println("Creating template and adding to cache...")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//
		log.Println("Using the cached template.")
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// Parse the templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add template to cache(map)
	templateCache[t] = tmpl

	return nil
}
