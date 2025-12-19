package constant

import (
	"testing"
)

func TestConstant(t *testing.T) {
	result := Constant("works")
	if result != "Constant works" {
		t.Error("Expected Constant to append 'works'")
	}
}
