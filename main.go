package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/employees", GetEmployees).Methods("GET")
	router.HandleFunc("/api/subscribe", Subscribe).Methods("POST")
	router.HandleFunc("/api/unsubscribe", Unsubscribe).Methods("POST")
	router.HandleFunc("/api/login", Login).Methods("POST")
	router.HandleFunc("/api/register", Register).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
