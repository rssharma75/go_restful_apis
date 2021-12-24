package main

import "net/http"

type server struct{}

type dog int

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (d dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dog"))
}

func main() {
	s := &server{}
	d := dog(10)

	http.Handle("/", s)
	http.Handle("/dog/", d)

	http.ListenAndServe(":7999", nil)
}
