package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: "test",
		Version:     "0.0.9",
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		log.Printf("Error: %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
