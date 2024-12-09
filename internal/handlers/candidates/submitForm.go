package candidates

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
	candidatesquery "github.com/Kunal-deve1oper/interview_app_backend/internal/services/candidatesQuery"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/utils"
)

func SubmitForm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	candidateInput := models.UserCandidate{}
	if err := json.NewDecoder(r.Body).Decode(&candidateInput); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format", "Failed to decode request body")
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	if !utils.ValidateCandidateData(candidateInput) {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing required fields", "Required fields are empty or invalid")
		log.Printf("Validation failed: missing required fields in request body: %v", candidateInput)
		return
	}

	err := candidatesquery.AddCandidateToDB(candidateInput)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Database operation failed", "Failed to add role to the database")
		log.Printf("Database insertion failed: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	jsonResponse := map[string]string{"msg": "created"}
	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
