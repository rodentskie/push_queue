package cursor

import (
	"testing"
)

func TestCursor(t *testing.T) {
	result := Cursor("works")
	if result != "Cursor works" {
		t.Error("Expected Cursor to append 'works'")
	}
}
