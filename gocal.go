package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cliffom/gocal/handlers"
)

func main() {
	port := getPort()

	log.Printf("Listening on 0.0.0.0:%s", port)
	http.HandleFunc("/api/classes/_healthcheck", handlers.Healthcheck)
	http.HandleFunc("/api/classes/", handlers.CreateClassListing)
	http.ListenAndServe(":"+port, nil)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return port
}
