package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cliffom/gocal/calendars"
	"github.com/cliffom/gocal/schema"
)

func createClass(w http.ResponseWriter, r *http.Request) {
	c := schema.Class{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error: %s", err)
		return
	}

	// Create class in whatever persistance layer we go with
	// Spawn job to create event
	go calendars.CreateEvent(&c)

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/classes/", createClass)
	http.ListenAndServe(":3000", nil)
}
