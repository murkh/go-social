package main

import "log"

func main() {
	cfg := Config{
		addr: "localhost:8080",
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
