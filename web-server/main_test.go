package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}

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
	req, _ := http.NewRequest("GET", "/foo", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusForbidden, res.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
