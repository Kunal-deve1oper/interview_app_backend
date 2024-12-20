package candidatesquery

import (
	"fmt"
	"log"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func AddCandidateToDB(data models.UserCandidate) error {
	query := `
			INSERT INTO "Candidates" 
			("id", "name", "email", "phoneNo", "photo", "gender", "country", "cv", "dob", "highestDegree", "highestDegreeOrg", "highestDegreeCGPA", "yog", "prevEmployer", "experience", "prevJobTitle", "duration", "isEmployed", "skills", "referralCode", "referralName", "links", "jobRole") 
			values 
			(gen_random_uuid(),$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22)
		`

	stmt, err := config.DB.Prepare(query)
	if err != nil {
		log.Printf("Failed to prepare query: %v", err)
		return fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Name, data.Email, data.PhoneNo, data.Photo, data.Gender, data.Country, data.Cv, data.Dob, data.HighestDegree, data.HighestDegreeOrg, data.HighestDegreeCGPA, data.Yog, data.PrevEmployer, data.Experience, data.PrevJobTitle, data.Duration, data.IsEmployed, data.Skills, data.ReferralCode, data.ReferralName, data.Links, data.JobRole)

	if err != nil {
		return err
	}

	return nil
}
