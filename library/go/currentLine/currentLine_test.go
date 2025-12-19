package currentline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentLine(t *testing.T) {
	line := CurrentLine()

	assert.Equal(t, line, 10)
}
