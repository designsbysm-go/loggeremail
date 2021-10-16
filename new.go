package timberemail

import (
	"io"
)

type options struct {
	from     string
	host     string
	password string
	port     int
	subject  string
	to       []string
}

func New(subject string, from string, password string, to []string, host string, port int) io.Writer {
	if from == "" || len(to) == 0 || host == "" || port == 0 {
		return nil
	}

	return options{
		from:     from,
		host:     host,
		password: password,
		port:     port,
		subject:  subject,
		to:       to,
	}
}
