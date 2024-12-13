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
# function to get all job role of an admin

	path = /allJobRole
	method = GET
	authentication = Bearer token

# RESPONSE

	if all good
	status : 200

[

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
	},

]

	if error
	{
		"error": error message,
	}
*/
func AllRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	// accessing claims set my the middleware
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*models.UserClaims)
	if !ok {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Unable to get claims", "Unable to get claims")
		log.Print("Unable to get claims")
		return
	}

	// Validate required fields
	if strings.TrimSpace(claims.UserID["id"]) == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "admin id is missing", "admin id is missing")
		log.Print("Validation failed: admin id is missing")
		return
	}

	// getting all job roles from DB for a particular user
	jobRoles, err := jobrolequery.AllRoleFromDB(claims.UserID["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Database operation failed", "Failed to find all role from the database")
		log.Printf("failed to get all roles of a current admin: %v", err)
		return
	}

	// sending the data to the user
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jobRoles); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
