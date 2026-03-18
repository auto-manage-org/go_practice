package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 1)
	expected := 2
	if result != expected {
		t.Errorf("Add(1, 1) = %d; want %d", result, expected)
	}
}
