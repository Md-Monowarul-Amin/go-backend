package main

import (
	"log"
	"net/http"
	"time"
)

type config struct {
	addr string
}

type application struct {
	config config
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux

}

func (app *application) run(mux *http.ServeMux) error {
	// Start the application logic here

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Starting server on %s", app.config.addr)
	return srv.ListenAndServe()
}
