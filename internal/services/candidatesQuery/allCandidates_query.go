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

func AllCandidatesFromDB(id, adminId string) ([]models.Candidate, error) {
	var jobCandidates []models.Candidate
	ctx := context.Background()

	// Attempt to retrieve data from Redis cache
	cachedData, err := config.RedisClient.Get(ctx, adminId+"_"+id).Result()
	if err == nil {
		if jsonErr := json.Unmarshal([]byte(cachedData), &jobCandidates); jsonErr == nil {
			return jobCandidates, nil
		}
		log.Println("Failed to unmarshal cached data, falling back to database query")
	}

	query := `SELECT 
			c."id", c."name", c."email", c."phoneNo", c."photo", c."gender", c."country", c."cv", c."dob", c."highestDegree", c."highestDegreeOrg", c."highestDegreeCGPA", c."yog", c."prevEmployer", c."experience", c."prevJobTitle", c."duration", c."isEmployed", c."skills", c."referralCode", c."referralName", c."links", c."jobRole", c."selected", c."createdAt", c."updatedAt" 
			FROM "Candidates" c 
			INNER JOIN 
			"Roles" r 
			ON c."jobRole" = r."id"
			WHERE r."createdBy" = $1 AND c."jobRole" = $2
	`

	stmt, err := config.DB.Prepare(query)
	if err != nil {
		log.Printf("Failed to prepare query: %v", err)
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(adminId, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}

	defer rows.Close()

	// extracting and appending each row
	for rows.Next() {
		var jobCandidate models.Candidate
		err := rows.Scan(&jobCandidate.Id, &jobCandidate.Name, &jobCandidate.Email, &jobCandidate.PhoneNo, &jobCandidate.Photo, &jobCandidate.Gender, &jobCandidate.Country, &jobCandidate.Cv, &jobCandidate.Dob, &jobCandidate.HighestDegree, &jobCandidate.HighestDegreeOrg, &jobCandidate.HighestDegreeCGPA, &jobCandidate.Yog, &jobCandidate.PrevEmployer, &jobCandidate.Experience, &jobCandidate.PrevJobTitle, &jobCandidate.Duration, &jobCandidate.IsEmployed, &jobCandidate.Skills, &jobCandidate.ReferralCode, &jobCandidate.ReferralName, &jobCandidate.Links, &jobCandidate.JobRole, &jobCandidate.Selected, &jobCandidate.CreatedAt, &jobCandidate.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}
		jobCandidates = append(jobCandidates, jobCandidate)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	// Cache the result in Redis
	dataJson, jsonErr := json.Marshal(jobCandidates)
	if jsonErr == nil {
		cacheErr := config.RedisClient.Set(ctx, adminId+"_"+id, dataJson, 5*time.Minute).Err()
		if cacheErr != nil {
			log.Println("failed to store in redis")
		}
	}

	return jobCandidates, nil
}
