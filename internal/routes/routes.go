package routes

import (
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/jobrole"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /addJobRole", jobrole.AddRole)
}
