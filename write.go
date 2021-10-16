package timberemail

import (
	"net/smtp"
)

var SendMail = smtp.SendMail

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
