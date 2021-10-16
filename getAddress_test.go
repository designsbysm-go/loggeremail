package timberemail

import (
	"fmt"
	"testing"
)

func TestShouldFormatAddress(t *testing.T) {
	host := "host"
	port := 25
	address := getAddress(host, port)

	if address != fmt.Sprintf("%s:%d", host, port) {
		t.Errorf("address should be host:port, got: %v", address)
	}
}
