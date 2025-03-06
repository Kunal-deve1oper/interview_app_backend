package models

import "time"

type Calendar struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Desc       string    `json:"desc"`
	Date       time.Time `json:"date"`
	Time       time.Time `json:"time"`
	IsComplete bool      `json:"isComplete"`
	Candidates []string  `json:"candidates"`
	// JobRoleId         string    `json:"jobRoleId"`
	CreatedByAdmin    string    `json:"createdByAdmin"`
	CreatedByAdminOrg string    `josn:"createdByAdminOrg"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type UserCalendar struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Desc       string    `json:"desc"`
	Date       time.Time `json:"date"`
	Time       time.Time `json:"time"`
	IsComplete bool      `json:"isComplete"`
	Candidates []string  `json:"candidates"`
	// JobRoleId         string    `json:"jobRoleId"`
	CreatedByAdmin    string `json:"createdByAdmin"`
	CreatedByAdminOrg string `josn:"createdByAdminOrg"`
}
