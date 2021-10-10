package loggeremail

import (
	"fmt"
)

func getAddress(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
