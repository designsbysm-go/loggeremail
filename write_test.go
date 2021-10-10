package loggeremail

import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
	"testing"
)

func TestShouldFormatAddress(t *testing.T) {
	host := "host"
	port := "port"
	address := getAddress(host, port)

	if address != fmt.Sprintf("%s:%s", host, port) {
		t.Errorf("address should be host:port, got: %v", address)
	}
}

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

func TestSingleToAuthShouldReturnString(t *testing.T) {
	exemplar := "to"

	toList := []string{exemplar}
	toString := getTo(toList)

	if toString != exemplar {
		t.Errorf("auth should be %s, got: %s", exemplar, toString)
	}
}

func TestMultipleToAuthShouldReturnFormattedString(t *testing.T) {
	toList := []string{"to", "to"}

	exemplar := strings.Join(toList, ";")
	toString := getTo(toList)

	if toString != exemplar {
		t.Errorf("auth should be %s, got: %s", exemplar, toString)
	}
}

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

	w := New(subject, "from", "", toList, "host", "port")
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

	w := New("", "from", "password", []string{"to"}, "host", "port")
	n, err := w.Write([]byte("test"))

	if n != 0 {
		t.Errorf("n should be 0, got: %d", n)
	}

	if err == nil {
		t.Errorf("err should not be nil")
	}
}
