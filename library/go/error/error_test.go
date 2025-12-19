package error

import (
	"testing"
)

func TestError(t *testing.T) {
	result := Error("works")
	if result != "Error works" {
		t.Error("Expected Error to append 'works'")
	}
}
