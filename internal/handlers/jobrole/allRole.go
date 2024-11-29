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
		log.Printf("Error querying from database: %v", err)
		http.Error(w, "Failed to querying from database", http.StatusInternalServerError)
		return
	}

	// sending the data to the user
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jobRoles); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}
