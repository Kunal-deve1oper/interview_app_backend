package jobrolequery

import (
	"fmt"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func AllRoleFromDB(id string) ([]models.Role, error) {

	// query to find all the roles created by a user
	query := `
		SELECT id,name,skills,experience,minATS,createdBy,createdAt,updatedAt 
		FROM Roles
		WHERE createdBy = $1
	`

	// querying DB
	rows, err := config.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}

	defer rows.Close()

	var jobRoles []models.Role

	// extracting and appending each row
	for rows.Next() {
		var jobRole models.Role
		err := rows.Scan(&jobRole.Id, &jobRole.Name, &jobRole.Skills, &jobRole.Experience, &jobRole.MinATS, &jobRole.CreatedBy, &jobRole.CreatedAt, &jobRole.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}
		jobRoles = append(jobRoles, jobRole)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return jobRoles, nil
}
