package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World - Gorilla Mux")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":5000", router))
}
