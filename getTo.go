package timberemail

import (
	"strings"
)

func getTo(to []string) string {
	return strings.Join(to, ";")
}
