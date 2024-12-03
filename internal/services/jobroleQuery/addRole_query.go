package jobrolequery

import (
	"time"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func AddRoleToDB(data models.UserRole, createdBy string) (models.Role, error) {
	var id string
	var createdAt, updatedAt time.Time
	addedData := models.Role{}

	// query to add job role to table Roles in the database
	query := `
			INSERT INTO "Roles" (name, skills, experience, minATS, createdBy) 
			values ($1, $2, $3, $4, $5) 
			RETURNING id,createdAt,updatedAt
		`

	// querying DB
	err := config.DB.QueryRow(query, data.Name, data.Skills, data.Experience, data.MinATS, createdBy).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return addedData, err
	}

	// Populate addedData with the inserted values
	addedData.Id = id
	addedData.Name = data.Name
	addedData.Skills = data.Skills
	addedData.Experience = data.Experience
	addedData.MinATS = data.MinATS
	addedData.CreatedBy = createdBy
	addedData.CreatedAt = createdAt
	addedData.UpdatedAt = updatedAt
	return addedData, nil
}
