package main

import (
	"flag"

	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "http port")
	flag.Parse()

	a := &app{
		router: mux.NewRouter(),
		z:      zorro.New(zorro.NewUUIDv4Generator(), zorro.NewInMemoryStorage()),
		port:   port,
	}

	a.start()
}
