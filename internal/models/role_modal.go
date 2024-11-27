package models

import "time"

type Role struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Skills     string    `json:"skills"`
	Experience int       `json:"experience"`
	MinATS     int       `json:"minATS"`
	CreatedBy  string    `json:"createdBy"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UserRole struct {
	Name       string `json:"name"`
	Skills     string `json:"skills"`
	Experience int    `json:"experience"`
	MinATS     int    `json:"minATS"`
	CreatedBy  string `json:"createdBy"`
}
