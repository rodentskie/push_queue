package mongo

import (
	"testing"
)

func TestMongo(t *testing.T) {
	result := Mongo("works")
	if result != "Mongo works" {
		t.Error("Expected Mongo to append 'works'")
	}
}
