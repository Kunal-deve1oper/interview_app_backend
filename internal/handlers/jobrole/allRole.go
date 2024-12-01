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
path = /allJobRole
method = GET
authentication = Bearer token
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
	if strings.TrimSpace(claims.Id) == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "admin id is missing", "admin id is missing")
		log.Print("Validation failed: admin id is missing")
		return
	}

	// getting all job roles from DB for a particular user
	jobRoles, err := jobrolequery.AllRoleFromDB(claims.Id)
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
