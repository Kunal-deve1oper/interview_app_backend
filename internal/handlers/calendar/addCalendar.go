package calendar

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/middleware"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
	calendarquery "github.com/Kunal-deve1oper/interview_app_backend/internal/services/calendarQuery"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/utils"
)

func AddCalendar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	userInput := models.UserCalendar{}

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
	if userInput.Title == "" || userInput.Desc == "" || userInput.CreatedByAdmin == "" || userInput.CreatedByAdminOrg == "" || claims.UserID["id"] == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing required fields", "Required fields are empty or invalid")
		log.Printf("Validation failed: missing required fields in request body: %v", userInput)
		return
	}

	// checking authentication
	if userInput.CreatedByAdmin != claims.UserID["id"] || userInput.CreatedByAdminOrg != claims.UserID["orgId"] {
		utils.SendErrorResponse(w, http.StatusForbidden, "Action not allowed", "Action not allowed")
		log.Printf("Authentication failed: Action not allowed: %v", userInput)
		return
	}

	// add the role to the database
	res, err := calendarquery.AddCalendarToDB(&userInput, claims.UserID["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Database operation failed", "Failed to add calendar to the database")
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
