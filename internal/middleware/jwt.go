package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/utils"
)

type contextKey string

const UserClaimsKey contextKey = "userClaims"

// JWTMiddleware validates the token and passes claims to the context
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.SendErrorResponse(w, http.StatusUnauthorized, "Authorization header missing", "Authorization header not found")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.SendErrorResponse(w, http.StatusUnauthorized, "Malformed token", "Malformed token")
			return
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			utils.SendErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("%v", err), fmt.Sprintf("%v", err))
			return
		}

		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
