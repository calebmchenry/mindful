package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/calebmchenry/mindful/web-server/auth"
	"github.com/joho/godotenv"
)

var s Server

func TestMain(m *testing.M) {
	godotenv.Load()
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
	req, _ := http.NewRequest("GET", "/api/v1/logout", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusForbidden, res.Code)
}

func TestApi_Authorized(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/logout", nil)
	token, _ := auth.CreateToken("john.doe")
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
