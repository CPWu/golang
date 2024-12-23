package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/cpwu/golang/pkg/config"
	"github.com/cpwu/golang/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	return mux
}
