package main

import (
	"log"

	"github.com/murkh/go-social/internal/env"
	"github.com/murkh/go-social/internal/store"
)

func main() {
	cfg := Config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
