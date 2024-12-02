package models

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	UserID map[string]string `json:"userId"`
	jwt.RegisteredClaims
}
