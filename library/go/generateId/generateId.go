package generateid

import (
	"crypto/rand"
	"fmt"
)

// GenerateID creates a 16-byte ID where the first byte is the type
// and the remaining 15 bytes are randomly generated
func GenerateID(typ int) ([]byte, error) {
	if typ < 0 || typ > 255 {
		return nil, fmt.Errorf("typ must be between 0 and 255, got %d", typ)
	}

	id := make([]byte, 16)
	id[0] = byte(typ)

	// Generate random bytes for the remaining 15 bytes
	if _, err := rand.Read(id[1:]); err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return id, nil
}
