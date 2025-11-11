package View

import "testing"

func TestAdd(t *testing.T) {
	result := 2 + 5
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}
