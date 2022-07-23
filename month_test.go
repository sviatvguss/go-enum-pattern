package go_enum_pattern

import "testing"

func TestMonth(t *testing.T) {
	m := June
	expected := "June"
	if got := m.String(); got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}
