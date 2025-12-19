package generateid

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	tests := []struct {
		name      string
		inputType int
		expectErr bool
	}{
		{"Valid type 0", 0, false},
		{"Valid type 255", 255, false},
		{"Valid type 128", 128, false},
		{"Negative type", -1, true},
		{"Large type", 256, true},
		{"Large type 1024", 1024, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenerateID(tt.inputType)

			// Check for expected errors
			if (err != nil) != tt.expectErr {
				t.Errorf("GenerateID(%d) error = %v, expectErr %v", tt.inputType, err, tt.expectErr)
				return
			}

			// Skip further validation if error was expected
			if tt.expectErr {
				return
			}

			// Validate the result length is 16 bytes
			if len(result) != 16 {
				t.Errorf("GenerateID(%d) length = %d; want 16", tt.inputType, len(result))
			}

			// Validate the first byte matches the input type
			if result[0] != byte(tt.inputType) {
				t.Errorf("GenerateID(%d) first byte = %d; want %d", tt.inputType, result[0], byte(tt.inputType))
			}
		})
	}

	t.Run("uniqueness", func(t *testing.T) {
		typ := 1
		id1, err := GenerateID(typ)
		if err != nil {
			t.Fatalf("GenerateID failed: %v", err)
		}

		id2, err := GenerateID(typ)
		if err != nil {
			t.Fatalf("GenerateID failed: %v", err)
		}

		// IDs should be different due to random bytes
		if string(id1) == string(id2) {
			t.Errorf("GenerateID produced duplicate IDs: %v", id1)
		}
	})

	t.Run("randomness", func(t *testing.T) {
		typ := 42
		id, err := GenerateID(typ)
		if err != nil {
			t.Fatalf("GenerateID failed: %v", err)
		}

		// Check that not all bytes after the first are zero (highly unlikely with random data)
		allZero := true
		for i := 1; i < len(id); i++ {
			if id[i] != 0 {
				allZero = false
				break
			}
		}

		if allZero {
			t.Errorf("GenerateID produced ID with all zero random bytes: %v", id)
		}
	})

}
