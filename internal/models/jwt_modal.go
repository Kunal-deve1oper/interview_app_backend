package models

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	OrgId string `json:"orgId"`
	jwt.RegisteredClaims
}
