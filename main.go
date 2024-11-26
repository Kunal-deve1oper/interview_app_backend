package main

import (
	"log"

	"github.com/Kunal-deve1oper/interview_app_backend/config"
	"github.com/Kunal-deve1oper/interview_app_backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.InitDB()
	defer config.DB.Close()

	s := server.New()
	if err := s.Start(); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
