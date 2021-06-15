package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/form3tech-oss/jwt-go"
)

type claims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func getAuthSecret() string {
	// TODO: get secret from env variables
	return "My Secret"
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

type contextKey string

func addUserContext(r *http.Request, c claims) *http.Request {
	key := contextKey("user")
	return r.WithContext(context.WithValue(r.Context(), key, c.User))
}
