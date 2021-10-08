package loggeremail

import (
	"fmt"
	"net/smtp"
	"strings"
)

func (o options) Write(p []byte) (n int, err error) {
	var auth smtp.Auth
	if o.password != "" {
		auth = smtp.PlainAuth("", o.from, o.password, o.host)
	}

	to := strings.Join(o.to, ";")
	body := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n", to, o.subject))
	body = append(body, p...)
	body = append(body, []byte("\r\n")...)

	err = smtp.SendMail(fmt.Sprintf("%s:%s", o.host, o.port), auth, o.from, o.to, body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}
