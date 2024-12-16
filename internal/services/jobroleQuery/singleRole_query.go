package jobrolequery

import (
	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func SingleRoleFromDB(jobId, amdinId string) (models.Role, error) {
	query := `
		SELECT "id","name","skills","experience","minATS","createdBy","expired","createdAt","updatedAt" 
		FROM "Roles"
		WHERE "createdBy" = $1 AND "id" = $2
	`

	var roleData models.Role

	err := config.DB.QueryRow(query, amdinId, jobId).Scan(&roleData.Id, &roleData.Name, &roleData.Skills, &roleData.Experience, &roleData.MinATS, &roleData.CreatedBy, &roleData.Expired, &roleData.CreatedAt, &roleData.UpdatedAt)

	if err != nil {
		return roleData, err
	}
	return roleData, nil
}
