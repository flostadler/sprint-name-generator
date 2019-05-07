package main

import (
	"fmt"
	"github.com/flostadler/name-generator/pkg/generators/docker"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/docker", GenerateDockerName).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func GenerateDockerName(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s", docker.Generate())
}
