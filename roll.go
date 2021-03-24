package main

import (
	"GO-GITHUB/view"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting server...")
	r := mux.NewRouter()
	r.HandleFunc("/register", view.RegisterView).Methods("POST")
	r.HandleFunc("/login", view.LoginView).Methods("POST")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
