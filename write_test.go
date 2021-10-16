package timberemail

import (
	"errors"
	"net/smtp"
	"testing"
)

func TestShouldSendEmail(t *testing.T) {
	old := SendMail
	defer func() { SendMail = old }()

	SendMail = func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return nil
	}

	to := "to"
	toList := []string{to}
	subject := "subject"
	p := []byte("test")

	w := New(subject, "from", "", toList, "host", 25)
	body := getBody(to, subject, p)
	n, err := w.Write(p)

	if n != len(body) {
		t.Errorf("n should be %d, got: %d", len(body), n)
	}

	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}
}

func TestShouldNotSendEmail(t *testing.T) {
	old := SendMail
	defer func() { SendMail = old }()

	SendMail = func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return errors.New("some error")
	}

	w := New("", "from", "password", []string{"to"}, "host", 25)
	n, err := w.Write([]byte("test"))

	if n != 0 {
		t.Errorf("n should be 0, got: %d", n)
	}

	if err == nil {
		t.Errorf("err should not be nil")
	}
}
