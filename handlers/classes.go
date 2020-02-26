package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cliffom/gocal/schema"
)

// CreateClassListing handler
func CreateClassListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c := schema.Class{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error: %s", err)
		return
	}

	if listing, err := c.CreateListing(); err != nil {
		data, _ := json.Marshal(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
	} else {
		data, _ := json.Marshal(listing)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
