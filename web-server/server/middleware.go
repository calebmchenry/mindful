package server

import (
	"fmt"
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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		c := claims{}
		keyFunc := func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(getAuthSecret()), nil
		}

		// Verify Token
		_, err := jwt.ParseWithClaims(tokenString, &c, keyFunc)

		if err == nil {
			r = addUserContext(r, c)
			next.ServeHTTP(w, r)
		} else {
			// Unauthorized
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
