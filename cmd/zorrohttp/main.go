package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodiez/zorro"
)

var r = mux.NewRouter()
var mk = zorro.New(zorro.NewUUIDv4Generator(), zorro.NewInMemoryStorage())
var port = flag.Int("port", 8080, "http port")

func main() {
	flag.Parse()

	r.HandleFunc("/mask/{id}", maskHandler)
	r.HandleFunc("/unmask/{mask}", unmaskHandler)

	log.Printf("Listening for connections on :%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))
}

func maskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mask := mk.Mask(vars["id"])

	log.Printf(fmt.Sprintf("200 - %s - %s", r.URL.String(), mask))
	fmt.Fprint(w, mask)

	return
}

func unmaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := mk.Unmask(vars["mask"])

	if !ok {
		log.Printf(fmt.Sprintf("404 - %s ", r.URL.String()))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Printf(fmt.Sprintf("200 - %s - %s", r.URL.String(), id))
	fmt.Fprintf(w, id)
	return
}
