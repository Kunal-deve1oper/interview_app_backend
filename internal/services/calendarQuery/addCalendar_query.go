package calendarquery

import (
	"time"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
	"github.com/lib/pq"
)

func AddCalendarToDB(userData *models.UserCalendar, adminId string) (models.Calendar, error) {
	var id string
	var createdAt, updatedAt time.Time
	addedData := models.Calendar{}

	query := `INSERT INTO "Calendar"
			("id","title","desc","date","time","isComplete","candidates","createdByAdmin","createdByAdminOrg") 
			VALUES (gen_random_uuid(),$1,$2,$3,$4,$5,$6,$7,$8) 
			RETURNING "id","createdAt","updatedAt"
		`

	err := config.DB.QueryRow(query, userData.Title, userData.Desc, userData.Date, userData.Time, false, pq.Array(userData.Candidates), userData.CreatedByAdmin, userData.CreatedByAdminOrg).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return addedData, err
	}

	addedData.Id = id
	addedData.Title = userData.Title
	addedData.Desc = userData.Desc
	addedData.Date = userData.Date
	addedData.Time = userData.Time
	addedData.IsComplete = false
	addedData.Candidates = userData.Candidates
	// addedData.JobRoleId = userData.JobRoleId
	addedData.CreatedByAdmin = userData.CreatedByAdmin
	addedData.CreatedByAdminOrg = userData.CreatedByAdminOrg
	addedData.CreatedAt = createdAt
	addedData.UpdatedAt = updatedAt

	return addedData, nil
}
