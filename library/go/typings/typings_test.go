package typings

import (
	"testing"
)

func TestTypings(t *testing.T) {
	result := Typings("works")
	if result != "Typings works" {
		t.Error("Expected Typings to append 'works'")
	}
}
