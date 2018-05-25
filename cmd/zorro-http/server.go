package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rodrigodiez/zorro/pkg/service"
)

type server struct {
	zorro  service.Zorro
	router *mux.Router
	port   int
}

func (s *server) maskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := s.zorro.Mask(vars["key"])

	w.Write([]byte(value))
	w.Header().Set("Content-Type", "text/plain")
}

func (s *server) unmaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, ok := s.zorro.Unmask(vars["value"])

	w.Header().Set("Content-Type", "text/plain")

	if ok != true {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))

		return
	}

	w.Write([]byte(key))
}

func (s *server) serve(debug bool) {
	log.Printf("Listening for connections on :%d\n", s.port)

	s.router.HandleFunc("/mask/{key}", s.maskHandler).Methods("POST")
	s.router.HandleFunc("/unmask/{value}", s.unmaskHandler).Methods("POST")

	if debug {
		log.Println("Expvars available in /debug/vars")
		s.router.Handle("/debug/vars", http.DefaultServeMux)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router))
}
