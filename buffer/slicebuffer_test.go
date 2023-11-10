package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrite(t *testing.T) {
	buf := NewSliceBuffer[byte]()

	buf.Write([]byte{1, 2, 3, 4})
	assert.Eqaul(t, 4, buf.Readable(T))
}
