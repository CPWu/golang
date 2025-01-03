package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cpwu/bookings/pkg/config"
	"github.com/cpwu/bookings/pkg/handlers"
	"github.com/cpwu/bookings/pkg/render"
)

// Concepts
// Using chi for routing

const portNumber string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// Production Flag
	app.InProduction = false

	// Intialize Sessions
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
