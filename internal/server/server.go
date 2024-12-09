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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow the frontend's origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true, // Allow credentials (Bearer token)
	})
	log.Println("Starting server on port 8080")
	handler := c.Handler(s.mux)
	return http.ListenAndServe(":8080", handler)
}
