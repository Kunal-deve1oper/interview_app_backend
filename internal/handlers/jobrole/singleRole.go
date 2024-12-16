package jobrole

import (
	"database/sql"
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
# function to delete a job role

	path = /singleJobRole?id=<role id>
	method = GET
	authentication = Bearer token

# RESPONSE

	if all good
	status : 200

	{
		"id": "3a4c0070-b6fd-47c2-8c26-047d395a34e9",
		"name": "IT Specialist",
		"skills": "HTML, CSS, JS, Ruby on rails",
		"experience": 1,
		"minATS": 80,
		"createdBy": "94625258-5132-4034-b2b6-83b5f48c31b9",
		"expired": true,
		"createdAt": "2024-12-12T14:07:07.54Z",
		"updatedAt": "2024-12-12T14:54:26.083Z"
	}

	if error
	{
		"error": error message,
	}
*/
func SingleRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	// accessing claims set my the middleware
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*models.UserClaims)
	if !ok {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "unable to get claims", "unable to get claims")
		log.Print("unable to get claims")
		return
	}

	// Validate required fields
	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" && strings.TrimSpace(claims.UserID["id"]) == "" {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "id is missing", "id is missing")
		log.Print("Validation failed: id is missing")
		return
	}

	// getting a single job role from DB for a particular user with given id
	res, err := jobrolequery.SingleRoleFromDB(id, claims.UserID["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendErrorResponse(w, http.StatusNotFound, "no job found or admin didn't created the role", "no job found or admin didn't created the role")
			log.Printf("no job found or admin didn't created the role: %s", id)
			return
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, "internal server error", "internal server error")
		log.Printf("internal server for id: %s", id)
		return
	}

	// sending the data to the user
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
