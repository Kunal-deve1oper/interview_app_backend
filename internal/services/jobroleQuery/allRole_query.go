package jobrolequery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func AllRoleFromDB(id string) ([]models.Role, error) {
	var jobRoles []models.Role
	ctx := context.Background()

	// Attempt to retrieve data from Redis cache
	cachedData, err := config.RedisClient.Get(ctx, id).Result()
	if err == nil {
		if jsonErr := json.Unmarshal([]byte(cachedData), &jobRoles); jsonErr == nil {
			return jobRoles, nil
		}
		log.Println("Failed to unmarshal cached data, falling back to database query")
	}

	// query to find all the roles created by a user
	query := `
		SELECT "id","name","skills","experience","minATS","createdBy","expired","createdAt","updatedAt" 
		FROM "Roles"
		WHERE "createdBy" = $1
	`

	// querying DB
	rows, err := config.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}

	defer rows.Close()

	// extracting and appending each row
	for rows.Next() {
		var jobRole models.Role
		err := rows.Scan(&jobRole.Id, &jobRole.Name, &jobRole.Skills, &jobRole.Experience, &jobRole.MinATS, &jobRole.CreatedBy, &jobRole.Expired, &jobRole.CreatedAt, &jobRole.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}
		jobRoles = append(jobRoles, jobRole)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	// Cache the result in Redis
	dataJson, jsonErr := json.Marshal(jobRoles)
	if jsonErr == nil {
		cacheErr := config.RedisClient.Set(ctx, id, dataJson, 5*time.Minute).Err()
		if cacheErr != nil {
			log.Println("failed to store in redis")
		}
	}

	return jobRoles, nil
}
