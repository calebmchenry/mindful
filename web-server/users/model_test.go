package users

import (
	"testing"
)

func TestSetPassword_DoesNotAllowEmptyPassword(t *testing.T) {
	user := UserModel{}
	password := ""
	err := user.setPassword(password)

	if err == nil {
		t.Errorf("expected hashed password to not match regular password")
	}
}

func TestSetPassword_Hashes(t *testing.T) {
	user := UserModel{}
	password := "password123"
	user.setPassword(password)

	if user.PasswordHash == password {
		t.Errorf("expected hashed password to not match regular password")
	}
}

func TestCheckPassword(t *testing.T) {
	user := UserModel{}
	password := "password123"
	otherPassword := "123password"
	user.setPassword(password)

	err := user.checkPassword(password)
	if err != nil {
		t.Errorf("expected checkPassword to produce no error when password matches")
	}

	err = user.checkPassword(otherPassword)
	if err == nil {
		t.Errorf("expected checkPassword to produce error when password does not match")
	}
}
