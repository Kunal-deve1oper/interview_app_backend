package candidatesquery

import (
	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func AddCandidateToDB(data models.UserCandidate) error {
	query := `
			INSERT INTO "Candidates" 
			("id", "name", "email", "phoneNo", "photo", "country", "cv", "dob", "highestDegree", "highestDegreeCGPA", "yog", "prevEmployer", "experience", "prevJobTitle", "duration", "isEmployed", "skills", "referralCode", "referralName", "links", "jobRole", "selected") 
			values 
			(gen_random_uuid(),$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21)
		`

	_, err := config.DB.Exec(query, data.Name, data.Email, data.PhoneNo, data.Photo, data.Country, data.Cv, data.Dob, data.HighestDegree, data.HighestDegreeCGPA, data.Yog, data.PrevEmployer, data.Experience, data.PrevJobTitle, data.Duration, data.IsEmployed, data.Skills, data.ReferralCode, data.ReferralName, data.Links, data.JobRole, false)

	if err != nil {
		return err
	}

	return nil
}
