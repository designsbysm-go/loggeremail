package loggeremail

import (
	"fmt"
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
