package main

import (
	"context"
	"net/http"

	"github.com/form3tech-oss/jwt-go"
)

func CorsOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(mchenryc): Get origin from env variable
		origin := "*"

		w.Header().Set("Access-Control-Allow-Origin", origin)
		next.ServeHTTP(w, r)
	})
}

type Claims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(mchenryc): get secret from env variable
		secret := "My Secret"
		tokenString := r.Header.Get("Authorization")
		claims := Claims{}
		_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err == nil {
			r.WithContext(context.WithValue(r.Context(), "user", claims.User))
			next.ServeHTTP(w, r)
		} else {
			// Unauthorized
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
