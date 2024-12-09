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

func ExpireRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*models.UserClaims)
	if !ok {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "unable to get claims", "unable to get claims")
		log.Print("unable to get claims")
		return
	}

	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" && strings.TrimSpace(claims.UserID["id"]) == "" {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "id is missing", "id is missing")
		log.Print("Validation failed: id is missing")
		return
	}

	res, err := jobrolequery.ExpireJobRoleInDB(id, claims.UserID["id"])
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

	if rowsAffected == 0 {
		utils.SendErrorResponse(w, http.StatusNotFound, "id not found in the DB", "id not found in the DB")
		log.Print("id not found in the DB")
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse := map[string]string{"id": id}
	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}

}
