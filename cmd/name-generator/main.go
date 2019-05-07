package main

import (
	"fmt"
	"github.com/flostadler/name-generator/pkg/generators/docker"
	"github.com/flostadler/name-generator/pkg/generators/marvel"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, docker.Generate())
	}).Methods("GET")

	router.HandleFunc("/marvel", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, marvel.Generate())
	}).Methods("GET")

    log.Fatal(http.ListenAndServe(":80", router))
}
