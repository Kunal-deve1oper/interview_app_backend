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
# function to activate a job role

	path = /activateJobRole?id=<item_to_activate id>
	method = PUT
	authentication = Bearer token

# RESPONSE

	if all good
	status : 200

	{
		"id" : deleted role id
	}

	if error
	{
		"error": error message,
	}
*/
func ActivateRole(w http.ResponseWriter, r *http.Request) {
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

	// updating expired in the database
	res, err := jobrolequery.ActivateRoleInDB(id, claims.UserID["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "failed to update row", "failed to update row")
		log.Printf("failed to update row: %s", id)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "failed to get rows affected", "failed to get rows affected")
		log.Print("failed to get rows affected")
		return
	}

	// checking if the item to be deleted is found or not
	if rowsAffected == 0 {
		utils.SendErrorResponse(w, http.StatusNotFound, "id not found in the DB", "id not found in the DB")
		log.Print("id not found in the DB")
		return
	}

	// sending back Id of activated job role
	w.WriteHeader(http.StatusOK)
	jsonResponse := map[string]string{"id": id}
	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
