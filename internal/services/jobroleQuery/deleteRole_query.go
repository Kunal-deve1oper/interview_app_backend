package jobrolequery

import (
	"database/sql"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
)

func DeleteRoleFromDB(id, adminId string) (sql.Result, error) {

	// query to delete a job role
	query := `
			DELETE FROM "Roles" 
			WHERE id = $1 and createdBy = $2
		`

	res, err := config.DB.Exec(query, id, adminId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
