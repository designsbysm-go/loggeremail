package loggeremail

import (
	"fmt"
)

func getAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
