package main

import (
	"log"

	"github.com/murkh/go-social/internal/env"
)

func main() {
	cfg := Config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
