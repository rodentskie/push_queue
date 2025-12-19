package converter

import (
	"testing"
)

func TestConverter(t *testing.T) {
	result := Converter("works")
	if result != "Converter works" {
		t.Error("Expected Converter to append 'works'")
	}
}
