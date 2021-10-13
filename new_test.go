package loggeremail

import (
	"testing"
)

func TestNewShouldReturnWriter(t *testing.T) {
	result := New("test", "from", "password", []string{"to"}, "localhost", 25)

	if result == nil {
		t.Errorf("should return valid io.Writer, got: nil")
	}
}

func TestNewShouldReturnNilWithoutFrom(t *testing.T) {
	result := New("", "", "", []string{"to"}, "host", 25)

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}

func TestNewShouldReturnNilWithoutTo(t *testing.T) {
	result := New("", "from", "", []string{}, "host", 25)

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}

func TestNewShouldReturnNilWithoutHost(t *testing.T) {
	result := New("", "from", "", []string{"to"}, "", 25)

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}

func TestNewShouldReturnNilWithoutPort(t *testing.T) {
	result := New("", "from", "", []string{"to"}, "host", 0)

	if result != nil {
		t.Errorf("should return nil, got: %v", result)
	}
}
