package routes

import (
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/admin"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /allAdmin", admin.AllAdmin)
}
