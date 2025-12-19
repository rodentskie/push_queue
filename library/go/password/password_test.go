package password

import (
	"testing"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "mySecret123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}
	if !VerifyPassword(hash, password) {
		t.Error("VerifyPassword should return true for correct password")
	}
	if VerifyPassword(hash, "wrongPassword") {
		t.Error("VerifyPassword should return false for incorrect password")
	}
}
