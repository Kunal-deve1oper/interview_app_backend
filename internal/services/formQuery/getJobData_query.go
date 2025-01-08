package formquery

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
)

func GetJobDataFromDB(id string) (string, string, int, bool, error) {
	ctx := context.Background()

	cachedData, err := config.RedisClient.Get(ctx, "form_query_"+id).Result()
	if err == nil {
		var data map[string]interface{}
		if jsonErr := json.Unmarshal([]byte(cachedData), &data); jsonErr == nil {
			return data["name"].(string), data["skills"].(string), int(data["experience"].(float64)), data["expired"].(bool), nil
		}
	}

	query := `
			SELECT "name", "skills", "experience", "expired" 
			FROM "Roles"
			WHERE "id" = $1
		`
	var name, skills string
	var experience int
	var expired bool

	err = config.DB.QueryRow(query, id).Scan(&name, &skills, &experience, &expired)
	if err != nil {
		return name, skills, experience, expired, err
	}

	data := map[string]interface{}{
		"name":       name,
		"skills":     skills,
		"experience": experience,
		"expired":    expired,
	}

	dataJson, jsonErr := json.Marshal(data)
	if jsonErr == nil {
		cacheErr := config.RedisClient.Set(ctx, "form_query_"+id, dataJson, 5*time.Minute).Err()
		if cacheErr != nil {
			log.Println("failed to store in redis")
		}
	}

	return name, skills, experience, expired, err
}
