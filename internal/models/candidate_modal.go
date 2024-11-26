package models

import "time"

type Candidate struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone_No     string    `json:"phone_no"`
	Address      string    `json:"address"`
	Organization string    `json:"organization"`
	Experience   int       `json:"experience"`
	Role         string    `json:"role"`
	Cv           string    `json:"cv"`
	Selected     bool      `json:"selected"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}