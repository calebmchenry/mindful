package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/calebmchenry/mindful/web-server/env"
	"github.com/form3tech-oss/jwt-go"
)

type claims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

// SignUp attempts to create a user account for an email and password
// If successful then a jwt will be returned
func SignUp(email, password, passwordRepeat string) (string, error) {
	emailT := strings.TrimSpace(email)
	passwordT := strings.TrimSpace(password)
	passwordRepeatT := strings.TrimSpace(passwordRepeat)
	if len(emailT) == 0 {
		return "", fmt.Errorf("email cannot be empty")
	}
	if len(passwordT) == 0 {
		return "", fmt.Errorf("password cannot be empty")
	}
	if passwordRepeatT != passwordRepeat {
		return "", fmt.Errorf("passwords must match")
	}
	// TODO(mchenryc): check that email is unique
	emailIsUnique := true
	if !emailIsUnique {
		return "", fmt.Errorf("account with that email has already been created")
	}
	// TODO(mchenryc): add user to the DB
	token, err := CreateToken(email)
	if err != nil {
		return "", fmt.Errorf("failed to create jwt")
	}
	return token, nil
}

func CreateToken(userId string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secret := env.GetAuthSecret()
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Login(email, password string) (string, error) {
	if len(strings.TrimSpace(email)) == 0 {
		// TODO(mchenryc): do better email validation
		return "", fmt.Errorf("invalid email")
	}
	if len(strings.TrimSpace(password)) == 0 {
		// TODO(mchenryc): do better password validation
		return "", fmt.Errorf("invalid password")
	}

	// TODO(mchenryc): validate from a database
	isValidUser := true
	if !isValidUser {
		return "", fmt.Errorf("failed login attempt")
	}
	token, err := CreateToken(email)
	if err != nil {
		return "", fmt.Errorf("failed to create jwt")
	}
	return token, nil
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		c := claims{}
		keyFunc := func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(env.GetAuthSecret()), nil
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
