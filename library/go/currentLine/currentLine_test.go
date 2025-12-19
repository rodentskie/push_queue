package currentline

import (
	"testing"
)

func TestCurrentLine(t *testing.T) {
	result := CurrentLine("works")
	if result != "CurrentLine works" {
		t.Error("Expected CurrentLine to append 'works'")
	}
}
