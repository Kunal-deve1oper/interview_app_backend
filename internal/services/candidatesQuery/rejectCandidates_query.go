package candidatesquery

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
)

func UpdateRejectedToDB(id, adminId string) (sql.Result, error) {
	query := `
			UPDATE "Candidates"
			SET "selected" = $1, "updatedAt" = CURRENT_TIMESTAMP 
			WHERE "id" = $2
  			AND EXISTS (
    			SELECT 1
    			FROM "Roles" r
    			WHERE r."id" = "Candidates"."jobRole" AND r."createdBy" = $3
  				)
		`
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		log.Printf("Failed to prepare query: %v", err)
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("reject", id, adminId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
