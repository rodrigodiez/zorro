package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro"
)

type app struct {
	port   int
	router *mux.Router
	z      zorro.Zorro
}

func (a *app) run() {
	log.Printf("Listening for connections on :%d", a.port)

	a.router.HandleFunc("/mask/{id}", a.requestID(a.logger(a.maskHandler))).Methods("POST")
	a.router.HandleFunc("/unmask/{mask}", a.requestID(a.logger(a.unmaskHandler))).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.router))
}
