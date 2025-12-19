package xlsxreader

import (
	"testing"
)

func TestReadXLSX(t *testing.T) {
	// This is a placeholder test. You can add a sample xlsx file and test parsing logic here.
	_, err := ReadXLSX("testdata/sample.xlsx")
	if err != nil {
		t.Logf("Expected error if file does not exist, got: %v", err)
	}
}
