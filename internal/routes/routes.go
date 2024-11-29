package routes

import (
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/jobrole"
)

func RegisterRoutes(mux *http.ServeMux, middleware func(http.Handler) http.Handler) {
	mux.Handle("GET /allJobRole", middleware(http.HandlerFunc(jobrole.AllRole)))
	mux.Handle("POST /addJobRole", middleware(http.HandlerFunc(jobrole.AddRole)))
	mux.Handle("POST /updateJobRole", middleware((http.HandlerFunc(jobrole.UpdateRole))))
}
