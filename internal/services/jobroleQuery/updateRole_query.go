package jobrolequery

import (
	"time"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func UpdateRoleToDB(updatedInput models.Role) (models.Role, error) {

	var updated_at time.Time

	// query to update a row in table Role
	query := `
			UPDATE Roles
			SET name = $1, skills = $2, experience = $3, minATS = $4, updatedAt = CURRENT_TIMESTAMP
			WHERE id = $5
			RETURNING updatedAt
		`

	// querying DB
	err := config.DB.QueryRow(query, updatedInput.Name, updatedInput.Skills, updatedInput.Experience, updatedInput.MinATS, updatedInput.Id).Scan(&updated_at)
	if err != nil {
		return updatedInput, err
	}

	updatedInput.UpdatedAt = updated_at
	return updatedInput, nil
}
