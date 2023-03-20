package bytebuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceCast(t *testing.T) {
	b := make([]byte, 8)
	u16 := SliceCast[byte, uint16](b)
	assert.EqualValues(t, 4, len(u16))
}
