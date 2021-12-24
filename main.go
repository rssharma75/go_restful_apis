package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	router := httprouter.New()
	router.GET("/hello", hello)
	log.Fatal(http.ListenAndServe(":5000", router))
}
