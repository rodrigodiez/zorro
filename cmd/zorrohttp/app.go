package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro/pkg/service"
)

type app struct {
	port   int
	router *mux.Router
	z      service.Zorro
}

func (a *app) run() {
	log.Printf("Listening for connections on :%d", a.port)

	a.router.HandleFunc("/mask/{id}", a.requestID(a.logger(a.maskHandler))).Methods("POST")
	a.router.HandleFunc("/unmask/{value}", a.requestID(a.logger(a.unmaskHandler))).Methods("POST")
	a.router.Handle("/debug/vars", http.DefaultServeMux)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.router))
}
