package xlsxreader

import (
	"testing"
)

func TestXlsxReader(t *testing.T) {
	result := XlsxReader("works")
	if result != "XlsxReader works" {
		t.Error("Expected XlsxReader to append 'works'")
	}
}
