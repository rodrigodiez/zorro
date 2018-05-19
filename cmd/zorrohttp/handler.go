package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *app) maskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	value := a.z.Mask(vars["id"])

	fmt.Fprint(w, value)

	return
}

func (a *app) unmaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := a.z.Unmask(vars["value"])

	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, id)

	return
}
