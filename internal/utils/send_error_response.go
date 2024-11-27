package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, statusCode int, userMessage, logMessage string) {
	w.WriteHeader(statusCode)
	response := map[string]string{
		"error": userMessage,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to send error response: %v, original error: %s", err, logMessage)
	}
}
