package main

import (
	"GO-GITHUB/view"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", view.RegisterView).Methods("POST")
	r.HandleFunc("/login", view.LoginView).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
