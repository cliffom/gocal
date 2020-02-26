package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/cliffom/gocal/schema"
)

// Status is the struct used for healthchecks
type Status struct {
	Environment string `json:"environment"`
	Status      string `json:"status"`
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	status := &Status{
		Environment: env,
		Status:      "healthy",
	}

	data, _ := json.Marshal(status)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func createClassListing(w http.ResponseWriter, r *http.Request) {
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

func main() {
	port := getPort()

	log.Printf("Listening on 0.0.0.0:%s", port)
	http.HandleFunc("/api/classes/_healthcheck", healthcheck)
	http.HandleFunc("/api/classes/", createClassListing)
	http.ListenAndServe(":"+port, nil)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return port
}
