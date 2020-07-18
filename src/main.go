package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	const port = ":8001"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Server up and running...")
	})
	router.HandleFunc("/api/v1/events", createEvent).Methods("POST")
	router.HandleFunc("/api/v1/events", getAllEvents).Methods("GET")
	log.Println("Server listening on port : ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
