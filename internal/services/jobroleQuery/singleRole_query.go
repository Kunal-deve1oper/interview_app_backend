package jobrolequery

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func SingleRoleFromDB(jobId, amdinId string) (models.Role, error) {
	var roleData models.Role
	ctx := context.Background()

	// Attempt to retrieve data from Redis cache
	cachedData, err := config.RedisClient.Get(ctx, "single_"+jobId).Result()
	if err == nil {
		if jsonErr := json.Unmarshal([]byte(cachedData), &roleData); jsonErr == nil {
			if roleData.CreatedBy == amdinId {
				return roleData, nil
			}
		}
		log.Println("Failed to unmarshal cached data, falling back to database query")
	}

	query := `
		SELECT "id","name","skills","experience","minATS","createdBy","expired","createdAt","updatedAt" 
		FROM "Roles"
		WHERE "createdBy" = $1 AND "id" = $2
	`

	err = config.DB.QueryRow(query, amdinId, jobId).Scan(&roleData.Id, &roleData.Name, &roleData.Skills, &roleData.Experience, &roleData.MinATS, &roleData.CreatedBy, &roleData.Expired, &roleData.CreatedAt, &roleData.UpdatedAt)

	if err != nil {
		return roleData, err
	}

	// Cache the result in Redis
	dataJson, jsonErr := json.Marshal(roleData)
	if jsonErr == nil {
		cacheErr := config.RedisClient.Set(ctx, "single_"+jobId, dataJson, 5*time.Minute).Err()
		if cacheErr != nil {
			log.Println("failed to store in redis")
		}
	}

	return roleData, nil
}
