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

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func (s *server) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home"))
}

func main() {
	s := &server{}
	d := dog(10)

	http.Handle("/", s)
	http.Handle("/dog/", d)

	http.HandleFunc("/home", s.home)

	http.ListenAndServe(":7999", nil)
}
