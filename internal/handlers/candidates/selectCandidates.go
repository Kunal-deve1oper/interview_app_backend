package candidates

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/middleware"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
	candidatesquery "github.com/Kunal-deve1oper/interview_app_backend/internal/services/candidatesQuery"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/utils"
)

/*
# function to get candidates of a job role

	path = /selectCandidate?id=<candidate id>
	method = PUT
	authentication = Bearer token

# RESPONSE

	if all good
	status : 200
	{
		"id" : <selected candidate id>
	}

	if error
	{
		"error": error message,
	}
*/
func SelectCandidates(w http.ResponseWriter, r *http.Request) {
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
	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" || strings.TrimSpace(claims.UserID["id"]) == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "id is missing", "id is missing")
		log.Print("Validation failed: id is missing")
		return
	}

	res, err := candidatesquery.UpdateSelectedToDB(id, claims.UserID["id"])
	if err != nil {
		log.Println(err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "failed to update candidate", "failed to update candidate")
		log.Printf("failed to update candidate: %s", id)
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
		utils.SendErrorResponse(w, http.StatusNotFound, "id not found in the DB and ensure admin owns the role and candidate exists", "id not found in the DB and ensure admin owns the role and candidate exists")
		log.Print("id not found in the DB and ensure admin owns the role and candidate exists")
		return
	}

	// sending back Id of deleted job role
	w.WriteHeader(http.StatusOK)
	jsonResponse := map[string]string{"id": id}
	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
