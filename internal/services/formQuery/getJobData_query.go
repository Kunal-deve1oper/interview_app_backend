package formquery

import "github.com/Kunal-deve1oper/interview_app_backend/config"

func GetJobDataFromDB(id string) (string, string, int, error) {
	query := `
			SELECT "name", "skills", "experience" 
			FROM "Roles"
			WHERE "id" = $1
		`
	var name, skills string
	var experience int

	err := config.DB.QueryRow(query, id).Scan(&name, &skills, &experience)

	return name, skills, experience, err
}
