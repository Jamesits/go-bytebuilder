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

func TestCarCdr(t *testing.T) {
	b := []byte{1, 2, 3, 4, 5}
	car, cdr := CarCdr[uint16](b)
	assert.EqualValues(t, 2<<8+1, *car) // TODO: make it work on both endianness
	assert.EqualValues(t, 3, len(cdr))
}
