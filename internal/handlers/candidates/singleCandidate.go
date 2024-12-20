package candidates

import (
	"database/sql"
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
# function to get a single candidate

	path = /singleCandidate?candidateId=<candidate id>&roleId=<role id>
	method = GET
	authentication = Bearer token

# RESPONSE

	if all good
	status = 200
	{
		"id": "78ddecb1-8f13-4857-9ed4-8e67a926a965",
		"name": "Kunal Ghosh",
		"email": "johndoe@example.com",
		"phoneNo": "1234567890",
		"photo": "https://example.com/photos/johndoe.jpg",
		"gender": "Male",
		"country": "United States",
		"cv": "https://example.com/cvs/johndoe.pdf",
		"dob": "1990-01-01",
		"highestDegree": "Master's in Computer Science",
		"highestDegreeOrg": "NSEC",
		"highestDegreeCGPA": 3.8,
		"yog": "2015",
		"prevEmployer": "Tech Corp",
		"experience": 7,
		"prevJobTitle": "Software Engineer",
		"duration": "4",
		"isEmployed": true,
		"skills": "JavaScript,React,Node.js",
		"referralCode": "REF12345",
		"referralName": "Jane Smith",
		"links": "https://linkedin.com/in/johndoe,https://github.com/johndoe",
		"jobRole": "36b60df7-869d-4eae-a227-9c9cde8d74cf",
		"selected": false,
		"createdAt": "2024-12-13T09:58:00.529Z",
		"updatedAt": "2024-12-13T09:58:00.529Z"
	}

	if error
	{
		"error": error message,
	}
*/
func SingleCandidate(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	// accessing claims set my the middleware
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*models.UserClaims)
	if !ok {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "unable to get claims", "unable to get claims")
		log.Print("unable to get claims")
		return
	}

	// Validate required fields
	candidateId := r.URL.Query().Get("candidateId")
	roleId := r.URL.Query().Get("roleId")
	if strings.TrimSpace(candidateId) == "" && strings.TrimSpace(roleId) == "" && strings.TrimSpace(claims.UserID["id"]) == "" {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "id is missing", "id is missing")
		log.Print("Validation failed: id is missing")
		return
	}

	// getting a single candidate role from DB for a particular user with given id
	res, err := candidatesquery.SingleCandidateFromDB(candidateId, roleId, claims.UserID["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendErrorResponse(w, http.StatusNotFound, "candidate not found", "candidate not found")
			log.Print("candidate not found")
			return
		}
		utils.SendErrorResponse(w, http.StatusNotFound, "error fetching candidate", "error fetching candidate")
		log.Print("error fetching candidate")
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
