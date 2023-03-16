package bytebuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SomeStruct struct {
	Value1 uint16
	// there is an implicit padding
	Value2 uint64
}

func TestMarshalUnmarshal(t *testing.T) {
	var err error

	s1 := SomeStruct{
		Value1: 100,
		Value2: 500,
	}

	b, err := Marshal(&s1)
	assert.NoError(t, err)

	var s2 SomeStruct
	err = Unmarshal(b, &s2)
	assert.NoError(t, err)

	assert.EqualValues(t, s1.Value1, s2.Value1)
	assert.EqualValues(t, s1.Value2, s2.Value2)
}
