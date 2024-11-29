package server

import (
	"log"
	"net/http"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/middleware"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/routes"
	"github.com/rs/cors"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}

	routes.RegisterRoutes(s.mux, middleware.JWTMiddleware)
	return s
}

func (s *Server) Start() error {
	log.Println("Starting server on port 8080")
	handler := cors.Default().Handler(s.mux)
	return http.ListenAndServe(":8080", handler)
}
