package loggeremail

import (
	"fmt"
	"net/smtp"
	"strings"
)

var SendMail = smtp.SendMail

func getAddress(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func getAuth(host string, user string, password string) smtp.Auth {
	if password == "" {
		return nil
	}

	return smtp.PlainAuth("", user, password, host)
}

func getBody(to string, subject string, p []byte) []byte {
	body := []byte(fmt.Sprintf("To: %s\r\n", to))

	if subject != "" {
		body = append(body, []byte(fmt.Sprintf("Subject: %s\r\n", subject))...)
	}

	body = append(body, []byte("\r\n")...)
	body = append(body, p...)
	body = append(body, []byte("\r\n")...)

	return body
}

func getTo(to []string) string {
	return strings.Join(to, ";")
}

func (o options) Write(p []byte) (n int, err error) {
	auth := getAuth(o.host, o.from, o.password)
	to := getTo(o.to)
	body := getBody(to, o.subject, p)
	address := getAddress(o.host, o.port)

	err = SendMail(address, auth, o.from, o.to, body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}
