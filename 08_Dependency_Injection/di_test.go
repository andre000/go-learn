package dependency

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Andre")

	received := buffer.String()
	expected := "Hello, Andre"

	assertString(t, received, expected)
}

func assertString(t *testing.T, received string, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ received %q expected %q", received, expected)
	}
}
