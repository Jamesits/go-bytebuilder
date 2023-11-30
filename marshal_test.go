package bytebuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type StructA struct {
	Value1 uint16 // 2 byte
	// implicit padding of 6 byte
	Value2 uint64 // 8 byte
}

type StructB struct {
	Value3 uint32 // 4 byte
}

type StructAB struct {
	StructA // 16 byte
	StructB // 4 byte
	// implicit padding of 4 byte
}

func TestMarshalUnmarshal(t *testing.T) {
	var err error

	s1 := StructA{
		Value1: 100,
		Value2: 500,
	}

	b, err := Marshal(&s1)
	assert.NoError(t, err)
	s1Size := int(unsafe.Sizeof(s1))
	assert.EqualValues(t, len(b), s1Size)

	var s2 StructA
	assert.NotEqual(t, s1, s2)
	err = Unmarshal(b, &s2)
	assert.NoError(t, err)

	assert.EqualValues(t, s1.Value1, s2.Value1)
	assert.EqualValues(t, s1.Value2, s2.Value2)
}

func TestUnmarshalDifferentSize(t *testing.T) {
	b := make([]byte, 128) // arbitrary large number
	var s StructA
	err := Unmarshal(b, &s)
	assert.ErrorIs(t, err, SizeMismatch)
}

func TestPartialUnmarshal(t *testing.T) {
	// StructAB should be a little larger than StructA + StructB because of a padding
	assert.LessOrEqual(t, uint64(unsafe.Sizeof(StructA{})+unsafe.Sizeof(StructB{})), uint64(unsafe.Sizeof(StructAB{})))

	s1 := StructAB{
		StructA: StructA{
			Value1: 100,
			Value2: 300,
		},
		StructB: StructB{
			Value3: 500,
		},
	}
	b, err := Marshal(&s1)
	assert.NoError(t, err)

	var sA StructA
	err = Unmarshal(b, &sA)
	assert.ErrorIs(t, err, SizeMismatch)
	assert.EqualValues(t, s1.Value1, sA.Value1)
	assert.EqualValues(t, s1.Value2, sA.Value2)

	var sB StructB
	err = Unmarshal(b[unsafe.Sizeof(sA):], &sB)
	assert.ErrorIs(t, err, SizeMismatch)
	assert.EqualValues(t, s1.Value3, sB.Value3)
}
