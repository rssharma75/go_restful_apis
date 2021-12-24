package main

import "net/http"

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	s := &server{}

	http.Handle("/", s)

	http.ListenAndServe(":7999", nil)
}
