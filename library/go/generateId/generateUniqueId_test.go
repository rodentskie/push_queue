package generateid

import (
	"crypto/sha256"
	"strconv"
	"testing"
)

func TestGenerateUniqueId(t *testing.T) {
	tests := []struct {
		name      string
		inputName string
		inputSeed int
	}{
		{"Basic test", "testName", 42},
		{"Empty name", "", 123},
		{"Zero seed", "example", 0},
		{"Special characters", "name!@#$%^&*", 99},
		{"Long name", "thisisaverylongnametotestthefunctionality", 1001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate expected hash
			input := tt.inputName + strconv.Itoa(tt.inputSeed)
			expectedHash := sha256.Sum256([]byte(input))
			expected := expectedHash[:]

			// Call the function
			result := GenerateUniqueId(tt.inputName, tt.inputSeed)

			// Validate the result
			if string(result) != string(expected) {
				t.Errorf("GenerateUniqueId(%q, %d) = %x; want %x", tt.inputName, tt.inputSeed, result, expected)
			}
		})
	}
}
