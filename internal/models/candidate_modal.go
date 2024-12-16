package models

import "time"

type Candidate struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	PhoneNo           string    `json:"phoneNo"`
	Photo             string    `json:"photo"`
	Gender            string    `json:"gender"`
	Country           string    `json:"country"`
	Cv                string    `json:"cv"`
	Dob               string    `json:"dob"`
	HighestDegree     string    `json:"highestDegree"`
	HighestDegreeCGPA float64   `json:"highestDegreeCGPA"`
	Yog               string    `json:"yog"`
	PrevEmployer      string    `json:"prevEmployer"`
	Experience        int       `json:"experience"`
	PrevJobTitle      string    `json:"prevJobTitle"`
	Duration          string    `json:"duration"`
	IsEmployed        bool      `json:"isEmployed"`
	Skills            string    `json:"skills"`
	ReferralCode      string    `json:"referralCode"`
	ReferralName      string    `json:"referralName"`
	Links             string    `json:"links"`
	JobRole           string    `json:"jobRole"`
	Selected          bool      `json:"selected"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type UserCandidate struct {
	Name              string  `json:"name"`
	Email             string  `json:"email"`
	PhoneNo           string  `json:"phoneNo"`
	Photo             string  `json:"photo"`
	Gender            string  `json:"gender"`
	Country           string  `json:"country"`
	Cv                string  `json:"cv"`
	Dob               string  `json:"dob"`
	HighestDegree     string  `json:"highestDegree"`
	HighestDegreeCGPA float64 `json:"highestDegreeCGPA"`
	Yog               string  `json:"yog"`
	PrevEmployer      string  `json:"prevEmployer"`
	Experience        int     `json:"experience"`
	PrevJobTitle      string  `json:"prevJobTitle"`
	Duration          string  `json:"duration"`
	IsEmployed        bool    `json:"isEmployed"`
	Skills            string  `json:"skills"`
	ReferralCode      string  `json:"referralCode"`
	ReferralName      string  `json:"referralName"`
	Links             string  `json:"links"`
	JobRole           string  `json:"jobRole"`
	RoleName          string  `json:"roleName"`
}
