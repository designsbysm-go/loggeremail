package loggeremail

import (
	"net/smtp"
)

func getAuth(host string, user string, password string) smtp.Auth {
	if password == "" {
		return nil
	}

	return smtp.PlainAuth("", user, password, host)
}
