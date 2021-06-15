package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

var s Server

func TestMain(m *testing.M) {
	s = New()
	code := m.Run()
	os.Exit(code)
}

func TestLogin(t *testing.T) {
	var jsonStr = []byte("fake")
	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
}

func TestApi_Unauthorized(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/foo", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusForbidden, res.Code)
}

func TestApi_Authorized(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/foo", nil)
	token, _ := createToken("john.doe")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %v", token))
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func createToken(userId string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secret := getAuthSecret()
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
