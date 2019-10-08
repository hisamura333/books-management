package main

import (
	"github.com/hisamura333/books-management/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Index).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}