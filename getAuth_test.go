package timberemail

import (
	"testing"
)

func TestShouldReturnAuth(t *testing.T) {
	auth := getAuth("", "", "password")

	if auth == nil {
		t.Errorf("auth should not be nil")
	}
}

func TestAuthShouldReturnNil(t *testing.T) {
	auth := getAuth("", "", "")

	if auth != nil {
		t.Errorf("auth should be nil, got: %v", auth)
	}
}
