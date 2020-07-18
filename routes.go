package main

import (
	"../go-rest-api-firestore/entity"
	"../go-rest-api-firestore/repository"
	"encoding/json"
	"net/http"
)

var (
	eventRepo repository.EventRepository = repository.NewEventRepository()
)

func getAllEvents(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	events, err := eventRepo.GetAllEvents()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"Error while retrieving posts"}`))
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(events)
}

func createEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	event := entity.Event{}
	err := json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"Error decoding data"}`))
	}

	_, err = eventRepo.CreateEvent(&event)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"Error creating event"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(event)
}
