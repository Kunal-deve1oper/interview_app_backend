package formquery

import "github.com/Kunal-deve1oper/interview_app_backend/config"

func GetJobDataFromDB(id string) (string, string, int, bool, error) {
	query := `
			SELECT "name", "skills", "experience", "expired" 
			FROM "Roles"
			WHERE "id" = $1
		`
	var name, skills string
	var experience int
	var expired bool

	err := config.DB.QueryRow(query, id).Scan(&name, &skills, &experience, &expired)

	return name, skills, experience, expired, err
}
