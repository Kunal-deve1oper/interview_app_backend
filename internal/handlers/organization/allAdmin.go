package organization

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/middleware"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
	organizationquery "github.com/Kunal-deve1oper/interview_app_backend/internal/services/organizationQuery"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/utils"
)

/*
# Function to get all admin to a specific organisation

	path = /allAdmin
	method = GET
	authentication = Bearer token

# RESPONSE

	if all good
	status = 200
	[
		{
			"id": "e19a40ef-a732-45e7-a848-e19364166619",
			"name": "TCS",
			"organisation": "36b60df7-869d-4eae-a227-9c9cde8d74cf",
			"position": "HR",
			"avatar": "https://randomimg.com",
			"createdAt": "2024-12-12T15:55:45.567Z",
			"updatedAt": "2024-12-12T15:55:45.567Z"
		}
	]

	if error
	{
		"error": error message,
	}
*/
func AllAdmin(w http.ResponseWriter, r *http.Request) {
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
		utils.SendErrorResponse(w, http.StatusInternalServerError, "id is missing", "id is missing")
		log.Print("Validation failed: id is missing")
		return
	}

	// fetching all admin to a specific organisation
	res, err := organizationquery.AllAdminFromDB(claims)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Database operation failed", "Failed to find all role from the database")
		log.Printf("failed to get all roles of a current admin: %v", err)
		return
	}

	// checking if no rows are found
	if res == nil {
		res = []models.Admin{}
	}

	// sending the data to the user
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
