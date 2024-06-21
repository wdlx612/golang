package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server started")

	// // http.Handler
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", WelcomeHandler).Methods("GET")

	log.Print(http.ListenAndServe(":9090", router))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Welcome to page..!"
	w.Write([]byte(msg))
}
