package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// JWTAuth reads the Authorization Header which expects a Bearer token, validates it using the
// `secret` string. It extracts `user_id` from the token and stores it in the request context
// that is passed to the next handler.
func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			uidFloat, ok := claims["user_id"].(float64)
			if !ok {
				http.Error(w, "Invalid token user_id", http.StatusUnauthorized)
				return
			}

			userId := int64(uidFloat)
			ctx := context.WithValue(r.Context(), "userID", userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
