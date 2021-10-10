package loggeremail

import (
	"testing"
)

func TestNewShouldReturnWriter(t *testing.T) {
	result := New("test", "from", "password", []string{"to"}, "localhost", "port")

	if result == nil {
		t.Errorf("should return valid io.Writer, got: nil")
	}
}

func TestNewShouldReturnNilWithoutFrom(t *testing.T) {
	result := New("", "", "", []string{"to"}, "host", "port")

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}

func TestNewShouldReturnNilWithoutTo(t *testing.T) {
	result := New("", "from", "", []string{}, "host", "port")

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}

func TestNewShouldReturnNilWithoutHost(t *testing.T) {
	result := New("", "from", "", []string{"to"}, "", "port")

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}

func TestNewShouldReturnNilWithoutPort(t *testing.T) {
	result := New("", "from", "", []string{"to"}, "host", "")

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}
