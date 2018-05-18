package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro/pkg/generator"
	"github.com/rodrigodiez/zorro/pkg/masker"
	"github.com/rodrigodiez/zorro/pkg/storage"
)

var r = mux.NewRouter()
var mk = masker.New(generator.NewUUIDv4(), storage.NewMem())

func main() {
	r.HandleFunc("/mask/{key}", MaskHandler)
	r.HandleFunc("/unmask/{value}", UnmaskHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func MaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	value := mk.Mask(vars["key"])

	fmt.Fprintf(w, "Masked '%s' as '%s'\n", vars["key"], value)
	return
}

func UnmaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key, ok := mk.Unmask(vars["value"])

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "'%s' could not be unmasked\n", vars["value"])

		return
	}

	fmt.Fprintf(w, "Unmasked '%s' as '%s'\n", vars["value"], key)
	return
}
