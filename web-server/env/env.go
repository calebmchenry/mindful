package env

import (
	"os"
	"strings"
)

func GetAuthSecret() string {
	return os.Getenv("AUTH_SECRET")
}

func GetMongoUser() string {
	return os.Getenv("MONGO_USER")
}

func GetMongoPassword() string {
	return os.Getenv("MONGO_PASSWORD")
}

func GetHttpPort() string {
	value := os.Getenv("HTTP_PORT")
	if len(strings.TrimSpace(value)) == 0 {
		return "8080"
	}
	return value
}

func GetHttpOrigin() string {
	value := os.Getenv("HTTP_ORIGIN")
	if len(strings.TrimSpace(value)) == 0 {
		return "*"
	}
	return value
}
