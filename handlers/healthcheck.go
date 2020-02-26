package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

// Status is the struct used for healthchecks
type HealthcheckStatus struct {
	Environment string `json:"environment"`
	Status      string `json:"status"`
}

// Healthcheck handler
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	status := &HealthcheckStatus{
		Environment: env,
		Status:      "healthy",
	}

	data, _ := json.Marshal(status)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
