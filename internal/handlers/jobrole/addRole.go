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
# function to add a role to database

	route = /addJobRole
	body = json
	Request = POST

	example = {
	  "name":type string,
	  "skills":type string,
	  "experience": type int,
	  "minATS": type int,
	}

# RESPONSE

	if all good
	status code : 201
	{
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

	if error
	{
		"error": error message,
	}
*/
func AddRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	userInput := models.UserRole{}

	// accessing claims set my the middleware
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*models.UserClaims)
	if !ok {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Unable to get claims", "Unable to get claims")
		return
	}

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format", "Failed to decode request body")
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	// Validate required fields
	if strings.TrimSpace(userInput.Name) == "" || strings.TrimSpace(userInput.Skills) == "" || strings.TrimSpace(claims.UserID["id"]) == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing required fields", "Required fields are empty or invalid")
		log.Printf("Validation failed: missing required fields in request body: %v", userInput)
		return
	}

	// add the role to the database
	res, err := jobrolequery.AddRoleToDB(userInput, claims.UserID["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Database operation failed", "Failed to add role to the database")
		log.Printf("Database insertion failed: %v", err)
		return
	}

	// Successfully added role, respond with the new role data
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
