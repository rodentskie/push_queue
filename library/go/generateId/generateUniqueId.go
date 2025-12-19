package generateid

import (
	"crypto/sha256"
	"strconv"
)

func GenerateUniqueId(name string, seed int) []byte {
	input := name + strconv.Itoa(seed)
	hash := sha256.Sum256([]byte(input))
	return hash[:]
}
