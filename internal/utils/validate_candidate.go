package utils

import (
	"strings"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func ValidateCandidateData(data models.UserCandidate) bool {
	if strings.TrimSpace(data.Name) == "" || strings.TrimSpace(data.Email) == "" || strings.TrimSpace(data.PhoneNo) == "" || strings.TrimSpace(data.Photo) == "" || strings.TrimSpace(data.Country) == "" || strings.TrimSpace(data.Cv) == "" || strings.TrimSpace(data.Dob) == "" || strings.TrimSpace(data.HighestDegree) == "" || strings.TrimSpace(data.Yog) == "" || strings.TrimSpace(data.Skills) == "" || strings.TrimSpace(data.JobRole) == "" || strings.TrimSpace(data.Gender) == "" {
		return false
	}
	return true
}
