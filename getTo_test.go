package timberemail

import (
	"strings"
	"testing"
)

func TestSingleToShouldReturnString(t *testing.T) {
	exemplar := "to"

	toList := []string{exemplar}
	toString := getTo(toList)

	if toString != exemplar {
		t.Errorf("auth should be %s, got: %s", exemplar, toString)
	}
}

func TestMultipleToShouldReturnFormattedString(t *testing.T) {
	toList := []string{"to", "to"}

	exemplar := strings.Join(toList, ";")
	toString := getTo(toList)

	if toString != exemplar {
		t.Errorf("auth should be %s, got: %s", exemplar, toString)
	}
}
