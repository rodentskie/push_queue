package generateid

import (
	"testing"
)

func TestGenerateId(t *testing.T) {
	result := GenerateId("works")
	if result != "GenerateId works" {
		t.Error("Expected GenerateId to append 'works'")
	}
}
