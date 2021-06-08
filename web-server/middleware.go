package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

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
		tokenString := extractToken(r)

		// Verify Token
		claims := Claims{}
		_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// TODO(mchenryc): get secret from env variable
			secret := "My Secret"
			return []byte(secret), nil
		})

		if err == nil {
			// Forward user with request
			r.WithContext(context.WithValue(r.Context(), "user", claims.User))
			next.ServeHTTP(w, r)
		} else {
			// Unauthorized
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
