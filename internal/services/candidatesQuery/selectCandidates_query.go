package candidatesquery

import (
	"database/sql"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
)

func UpdateSelectedToDB(id, adminId string) (sql.Result, error) {
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

	res, err := config.DB.Exec(query, true, id, adminId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
