package candidates

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
	candidatesquery "github.com/Kunal-deve1oper/interview_app_backend/internal/services/candidatesQuery"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/utils"
)

/*
# function to add candidate

	path = /submitForm
	method = POST
	body = json

	data = {
		"name": "Rajit Dutta",
		"email": "johndoe@example.com",
		"phoneNo": "1234567890",
		"gender": "Male",
		"photo": "https://example.com/photos/johndoe.jpg",
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
		"links":"https://linkedin.com/in/johndoe,https://github.com/johndoe",
		"jobRole": "8c313455-83c1-4495-b5a5-d10a89225a31",
		"roleName" : "Go developer"
	}

# RESPONSE

	if all good
	status: 201
	{
		"msg": "created"
	}

	if error
	{
		"error": error message,
	}
*/
func SubmitForm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	candidateInput := models.UserCandidate{}
	if err := json.NewDecoder(r.Body).Decode(&candidateInput); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format", "Failed to decode request body")
		log.Printf("Failed to decode request body: %v", err)
		return
	}

	if !utils.ValidateCandidateData(candidateInput) {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Missing required fields", "Required fields are empty or invalid")
		log.Printf("Validation failed: missing required fields in request body: %v", candidateInput)
		return
	}

	err := candidatesquery.AddCandidateToDB(candidateInput)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Database operation failed", "Failed to add role to the database")
		log.Printf("Database insertion failed: %v", err)
		return
	}

	// mailBody := utils.SubmitMailTemplate(candidateInput.Name, candidateInput.RoleName, "Planet Interview", "xyz", candidateInput.Email, "support@planetinterview.com")
	// utils.SendMail("Application Confirmation", candidateInput.Email, mailBody)
	w.WriteHeader(http.StatusCreated)
	jsonResponse := map[string]string{"msg": "created"}
	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Response generation failed", "Failed to encode JSON response")
		return
	}
}
