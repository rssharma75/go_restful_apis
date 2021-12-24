package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

type server struct {
	user User
}

type User struct {
	UserName string `json:username`
	Email    string `json:email`
	Age      int    `json:age`
}

type dog int

func getBase64(w http.ResponseWriter, r *http.Request) {
	msg := strings.Split(r.URL.String(), "/")[2]
	data := []byte(msg)
	str := base64.StdEncoding.EncodeToString(data)
	w.Write([]byte(str))
}
func (s *server) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	switch r.Method {
	case "GET":
		e := json.NewEncoder(w)
		e.Encode(s.user)
		w.WriteHeader(http.StatusOK)
	case "PUT":
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.user = user
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"update":"OK"}`))

	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

}
func main() {
	s := &server{
		user: User{
			UserName: "go_learner",
			Email:    "go.learner@gmail.com",
			Age:      10,
		},
	}
	http.HandleFunc("/user", s.getUser)
	http.HandleFunc("/base64/", getBase64)

	http.ListenAndServe(":5000", nil)
}
