package auth

import (
	"testing"
)

func TestLogin(t *testing.T) {
	email := "calebmchenry@gmail.com"
	password := "pass1234"
	emptyString := "  "
	// empty email
	_, err := Login(emptyString, password)
	if err == nil {
		t.Errorf("expected empty email to produce an error")
	}

	// empty password
	_, err = Login(email, emptyString)
	if err == nil {
		t.Errorf("expected empty password to produce an error")
	}

	// TODO(mchenryc): authenticate creds

	// valid
	token, err := Login(email, password)
	if len(token) == 0 {
		t.Errorf("expected valid inputs to produce a jwt")
	}
	if err != nil {
		t.Errorf("expected valid creds to not produce an error")
	}
}

func TextSignUp(t *testing.T) {
	validEmail := "calebmchenry@gmail.com"
	validPassword := "pass1234"
	otherPassword := "pass1234"
	emptyString := "  "
	// empty email
	_, err := SignUp(emptyString, validPassword, validPassword)
	if err == nil {
		t.Errorf("expected empty email to produce an error")
	}
	// TODO(mchenryc): invalid email

	// empty pass
	_, err = SignUp(validEmail, emptyString, validPassword)
	if err == nil {
		t.Errorf("expected empty password to produce an error")
	}

	// not matching pass
	_, err = SignUp(validEmail, validPassword, otherPassword)
	if err == nil {
		t.Errorf("expected mismatched passwords to produce an error")
	}

	// TODO(mchenry): not unique email

	// valid
	token, err := SignUp(validEmail, validPassword, validPassword)
	if len(token) == 0 {
		t.Errorf("expected valid inputs to produce a jwt")
	}
	if err != nil {
		t.Errorf("expected valid inputs to produce no error")
	}
}
