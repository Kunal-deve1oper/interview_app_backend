package jobrole

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/middleware"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
	jobrolequery "github.com/Kunal-deve1oper/interview_app_backend/internal/services/jobroleQuery"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/utils"
)

/*
path = /updateJobRole
method = POST
authentication = Bearer token

	example payload = {
	  "id": "48e160e1-3a66-4dc5-a9ee-aba701fc1fb5",
	  "name": "Go developer",
	  "skills": "GO,CSS,HTML,JAVASCRIPT,AWS",
	  "experience": 3,
	  "minATS": 80,
	  "createdBy": "a6316878-b270-4b41-9d29-647e0478d1e3",
	  "expired": false,
	  "createdAt": "2024-11-29T19:48:49.925669Z",
	  "updatedAt": "2024-11-29T19:48:49.925669Z"
	}
*/
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	updatedInput := models.Role{}

	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*models.UserClaims)
	if !ok {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "unable to get claims", "unable to get claims")
		log.Print("unable to get claims")
		return
	}

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&updatedInput); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format", "Failed to decode request body")
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	// Validate required fields
	if strings.TrimSpace(updatedInput.Name) == "" || strings.TrimSpace(updatedInput.Skills) == "" || strings.TrimSpace(updatedInput.Id) == "" || strings.TrimSpace(updatedInput.CreatedBy) == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing required fields", "Required fields are empty or invalid")
		log.Printf("Validation failed: missing required fields in request body: %v", updatedInput)
		return
	}

	// update contents to DB
	res, err := jobrolequery.UpdateRoleToDB(updatedInput, claims.UserID["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Database operation failed", "Failed to update role to the database")
		log.Printf("Database role updation failed: %v", err)
		return
	}

	// send updated data
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
