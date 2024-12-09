package form

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	formquery "github.com/Kunal-deve1oper/interview_app_backend/internal/services/formQuery"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/utils"
)

func GetFormData(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "id is missing", "id is missing")
		log.Print("Validation failed: id is missing")
		return
	}

	name, skills, experience, err := formquery.GetJobDataFromDB(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendErrorResponse(w, http.StatusNotFound, "no job found", "no job found")
			log.Printf("no job found with id: %s", id)
			return
		}
		utils.SendErrorResponse(w, http.StatusInternalServerError, "internal server error", "internal server error")
		log.Printf("internal server for id: %s", id)
		return
	}

	jsonResponse := map[string]interface{}{
		"name":       name,
		"skills":     skills,
		"experience": experience,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
