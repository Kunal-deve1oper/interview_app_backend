package jobrolequery

import (
	"database/sql"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
)

func ExpireJobRoleInDB(id, adminId string) (sql.Result, error) {

	query := `
			UPDATE "Roles"
			SET "expired" = $1, "updatedAt" = CURRENT_TIMESTAMP
			WHERE "id" = $2 and "createdBy" = $3
		`

	res, err := config.DB.Exec(query, true, id, adminId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
