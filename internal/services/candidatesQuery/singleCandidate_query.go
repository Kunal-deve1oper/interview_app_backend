package candidatesquery

import (
	"fmt"
	"log"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func SingleCandidateFromDB(candidateId, roleId, adminId string) (*models.Candidate, error) {
	query := `SELECT 
			c."id", c."name", c."email", c."phoneNo", c."photo", c."gender", c."country", c."cv", c."dob", c."highestDegree", c."highestDegreeOrg", c."highestDegreeCGPA", c."yog", c."prevEmployer", c."experience", c."prevJobTitle", c."duration", c."isEmployed", c."skills", c."referralCode", c."referralName", c."links", c."jobRole", c."selected", c."createdAt", c."updatedAt" 
			FROM "Candidates" c 
			INNER JOIN 
			"Roles" r 
			ON c."jobRole" = r."id"
			WHERE r."createdBy" = $1 AND c."jobRole" = $2 AND c."id" = $3
	`

	stmt, err := config.DB.Prepare(query)
	if err != nil {
		log.Printf("Failed to prepare query: %v", err)
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var jobCandidate models.Candidate

	err = stmt.QueryRow(adminId, roleId, candidateId).Scan(
		&jobCandidate.Id,
		&jobCandidate.Name,
		&jobCandidate.Email,
		&jobCandidate.PhoneNo,
		&jobCandidate.Photo,
		&jobCandidate.Gender,
		&jobCandidate.Country,
		&jobCandidate.Cv,
		&jobCandidate.Dob,
		&jobCandidate.HighestDegree,
		&jobCandidate.HighestDegreeOrg,
		&jobCandidate.HighestDegreeCGPA,
		&jobCandidate.Yog,
		&jobCandidate.PrevEmployer,
		&jobCandidate.Experience,
		&jobCandidate.PrevJobTitle,
		&jobCandidate.Duration,
		&jobCandidate.IsEmployed,
		&jobCandidate.Skills,
		&jobCandidate.ReferralCode,
		&jobCandidate.ReferralName,
		&jobCandidate.Links,
		&jobCandidate.JobRole,
		&jobCandidate.Selected,
		&jobCandidate.CreatedAt,
		&jobCandidate.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &jobCandidate, nil
}
