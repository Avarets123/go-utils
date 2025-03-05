package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type BadAuth struct {
	Username, Password string
}

func (b *BadAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if username != b.Username && password != b.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	newCtx := context.WithValue(r.Context(), "username", username)
	r = r.WithContext(newCtx)
	next(w, r)

}

func hello(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username")
	fmt.Fprintf(w, "Hi %s", username)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")

	n := negroni.Classic()

	n.Use(&BadAuth{
		Username: "admin",
		Password: "123",
	})

	n.UseHandler(r)

	http.ListenAndServe(":3001", n)

}
