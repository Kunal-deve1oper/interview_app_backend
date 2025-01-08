package candidatesquery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func SingleCandidateFromDB(candidateId, roleId, adminId string) (*models.Candidate, error) {
	var jobCandidate models.Candidate
	ctx := context.Background()

	// Attempt to retrieve data from Redis cache
	cachedData, err := config.RedisClient.Get(ctx, candidateId+"_"+adminId).Result()
	if err == nil {
		if jsonErr := json.Unmarshal([]byte(cachedData), &jobCandidate); jsonErr == nil {
			if jobCandidate.JobRole == roleId {
				return &jobCandidate, nil
			}
		}
		log.Println("Failed to unmarshal cached data, falling back to database query")
	}

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

	// Cache the result in Redis
	dataJson, jsonErr := json.Marshal(jobCandidate)
	if jsonErr == nil {
		cacheErr := config.RedisClient.Set(ctx, candidateId+"_"+adminId, dataJson, 5*time.Minute).Err()
		if cacheErr != nil {
			log.Println("failed to store in redis")
		}
	}
	return &jobCandidate, nil
}
