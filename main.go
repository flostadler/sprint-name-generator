package main

import (
	"fmt"
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/docker/docker/pkg/namesgenerator"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/generate", GenerateName).Methods("GET")
    log.Fatal(http.ListenAndServe(":80", router))
}

func GenerateName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", namesgenerator.GetRandomName(0))
}
