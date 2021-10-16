package timberemail

import (
	"fmt"
)

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
