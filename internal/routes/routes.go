package routes

import (
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/candidates"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/form"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/jobrole"
)

func RegisterRoutes(mux *http.ServeMux, middleware func(http.Handler) http.Handler) {
	mux.Handle("GET /allJobRole", middleware(http.HandlerFunc(jobrole.AllRole)))
	mux.Handle("POST /addJobRole", middleware(http.HandlerFunc(jobrole.AddRole)))
	mux.Handle("PUT /updateJobRole", middleware(http.HandlerFunc(jobrole.UpdateRole)))
	mux.Handle("PUT /expireJobRole", middleware(http.HandlerFunc(jobrole.ExpireRole)))
	mux.Handle("DELETE /deleteJobRole", middleware(http.HandlerFunc(jobrole.DeleteRole)))

	mux.HandleFunc("GET /formJobData", form.GetFormData)
	mux.HandleFunc("POST /submitForm", candidates.SubmitForm)
}
