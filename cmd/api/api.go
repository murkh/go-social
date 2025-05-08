package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/murkh/go-social/internal/store"
)

type application struct {
	config Config
	store  store.Storage
}

type Config struct {
	addr string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})
	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
	}

	log.Default().Println("Starting server on", app.config.addr)

	return srv.ListenAndServe()
}
