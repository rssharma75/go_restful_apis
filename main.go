package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World - Chi")
}

func main() {
	router := chi.NewRouter()
	router.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":5000", router))
}
