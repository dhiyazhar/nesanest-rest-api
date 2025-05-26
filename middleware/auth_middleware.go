package middleware

import (
    "nesanest-rest-api/helper"
    "nesanest-rest-api/model/web"
    "net/http"
    "strings"
)

type AuthMiddleware struct {
    Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
    return &AuthMiddleware{
        Handler: handler,
    }
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    if request.Header.Get("X-API-Key") == "RAHASIA" {
        middleware.Handler.ServeHTTP(writer, request)
    } else {
        writer.Header().Set("Content-Type", "application/json")
        writer.WriteHeader(http.StatusUnauthorized)
        webResponse := web.WebResponse{
            Code:   http.StatusUnauthorized,
            Status: "UNAUTHORIZED",
        }
        helper.WriteToResponseBody(writer, webResponse)
    }
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Unauthorized: missing or invalid token", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := helper.ParseJWT(tokenString)
        if err != nil {
            http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
            return
        }

		_ = claims

        next.ServeHTTP(w, r)
    })
}