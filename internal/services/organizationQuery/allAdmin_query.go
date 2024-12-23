package organizationquery

import (
	"fmt"
	"log"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
)

func AllAdminFromDB(claims *models.UserClaims) ([]models.Admin, error) {
	query := `
			SELECT 
			a."id", a."name", a."email", a."organisation", a."position", a."avatar", a."createdAt", a."updatedAt" 
			FROM 
			"Admins" a 
			INNER JOIN 
			"Organisations" o 
			ON 
			a."organisation" = o."id"
			WHERE o."id" = $1
		`

	stmt, err := config.DB.Prepare(query)
	if err != nil {
		log.Printf("Failed to prepare query: %v", err)
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(claims.UserID["id"])
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}

	defer rows.Close()

	var allAdmin []models.Admin

	for rows.Next() {
		var admin models.Admin
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Email, &admin.Organization, &admin.Position, &admin.Avatar, &admin.CreatedAt, &admin.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}
		allAdmin = append(allAdmin, admin)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return allAdmin, nil
}
