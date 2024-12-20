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

	path = /allCandidate?id=<job_role id>
	method = GET
	authentication = Bearer token

# RESPONSE

	if all good and rows found
	status : 200
	[
		{
			"id": "e19a40ef-a732-45e7-a848-e19364166619",
			"name": "John Doe",
			"email": "random@gmail.com",
			"phoneNo": "+91 7112343456",
			"photo": "https://qjjltpmoudplzbqwjbvy.supabase.co/storage/v1/object/public/candidate-photo/Z-or547uFsyte1Qi-U85p.jpeg",
			"gender": "Rajit",
			"country": "India",
			"cv": "https://qjjltpmoudplzbqwjbvy.supabase.co/storage/v1/object/public/candidate-cv/VB2fmDw6p4kGyKTDfgavi.pdf",
			"dob": "2024-12-04",
			"highestDegree": "B.Tech in CSE",
			"highestDegreeOrg": "NSEC",
			"highestDegreeCGPA": 9,
			"yog": "2024",
			"prevEmployer": "",
			"experience": 0,
			"prevJobTitle": "",
			"duration": "",
			"isEmployed": false,
			"skills": "go,ceo",
			"referralCode": "",
			"referralName": "",
			"links": "hfgfafaf",
			"jobRole": "36b60df7-869d-4eae-a227-9c9cde8d74cf",
			"selected": false,
			"createdAt": "2024-12-12T15:55:45.567Z",
			"updatedAt": "2024-12-12T15:55:45.567Z"
		}
	]

	if error
	{
		"error": error message,
	}
*/
func AllCandidates(w http.ResponseWriter, r *http.Request) {
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

	// fetching all candidate to a specific role
	res, err := candidatesquery.AllCandidatesFromDB(id, claims.UserID["id"])
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Database operation failed", "Failed to find all role from the database")
		log.Printf("failed to get all roles of a current admin: %v", err)
		return
	}

	// checking if no rows are found
	if res == nil {
		res = []models.Candidate{}
	}

	// sending the data to the user
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
