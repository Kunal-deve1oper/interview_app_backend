package routes

import (
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/candidates"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/form"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/handlers/jobrole"
)

func RegisterRoutes(mux *http.ServeMux, middleware func(http.Handler) http.Handler) {
	// job role routes
	mux.Handle("GET /allJobRole", middleware(http.HandlerFunc(jobrole.AllRole)))
	mux.Handle("GET /singleJobRole", middleware(http.HandlerFunc(jobrole.SingleRole)))
	mux.Handle("POST /addJobRole", middleware(http.HandlerFunc(jobrole.AddRole)))
	mux.Handle("PUT /updateJobRole", middleware(http.HandlerFunc(jobrole.UpdateRole)))
	mux.Handle("PUT /expireJobRole", middleware(http.HandlerFunc(jobrole.ExpireRole)))
	mux.Handle("PUT /activateJobRole", middleware(http.HandlerFunc(jobrole.ActivateRole)))
	mux.Handle("DELETE /deleteJobRole", middleware(http.HandlerFunc(jobrole.DeleteRole)))

	// getting form data route
	mux.HandleFunc("GET /formJobData", form.GetFormData)

	// candidates routes
	mux.Handle("GET /allCandidate", middleware(http.HandlerFunc(candidates.AllCandidates)))
	mux.Handle("GET /singleCandidate", middleware(http.HandlerFunc(candidates.SingleCandidate)))
	mux.HandleFunc("POST /submitForm", candidates.SubmitForm)
	mux.Handle("PUT /selectCandidate", middleware(http.HandlerFunc(candidates.SelectCandidates)))
	mux.Handle("PUT /rejectCandidate", middleware(http.HandlerFunc(candidates.RejectCandidate)))
}
