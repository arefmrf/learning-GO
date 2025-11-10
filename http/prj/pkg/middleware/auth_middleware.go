package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"prj/internal/modules/user/service"
	"strings"
	"time"
)

type contextKey string

const UserIDKey = contextKey("user_id")

func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			claims := &service.CustomClaims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			// 1 Check parsing or signature errors
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// 2 Validate token + claims
			if !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// 3 Optional manual expiration check (library already does this)
			if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}

			fmt.Println("✅ Authenticated user:", claims.Username)

			// 4 Add user info to context
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
