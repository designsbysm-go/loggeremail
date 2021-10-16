package timberemail

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBodyShouldBeFormatted(t *testing.T) {
	to := "to"
	subject := "subject"
	p := []byte("message")

	body := getBody(to, subject, p)

	if !bytes.Contains(body, []byte(fmt.Sprintf("To: %s", to))) {
		t.Errorf("body should contain to, got: %s", string(body))
	}

	if !bytes.Contains(body, []byte(fmt.Sprintf("Subject: %s", subject))) {
		t.Errorf("body should contain subject, got: %s", string(body))
	}

	if !bytes.Contains(body, p) {
		t.Errorf("body should contain p, got: %s", string(body))
	}
}

func TestBodyShouldNotContainSubject(t *testing.T) {
	body := getBody("", "", []byte{})

	if bytes.Contains(body, []byte("Subject: ")) {
		t.Errorf("body should not contain subject, got: %s", string(body))
	}
}
