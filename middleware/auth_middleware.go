package middleware

import (
	"context"
	"nesanest-rest-api/helper"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := ""
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			http.Error(w, "Bearer Token Required", http.StatusUnauthorized)
			return
		}

		claims, err := helper.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", int(claims["user_id"].(float64)))
		ctx = context.WithValue(ctx, "email", claims["email"].(string))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
