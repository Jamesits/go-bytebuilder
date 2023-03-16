package bytebuilder

import (
	"bytes"
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type Victim struct {
	A bool   // 1 byte, padded to 8 bytes
	B uint64 // 8 bytes
	C bool   // 1 byte
}

func TestWriteObject(t *testing.T) {
	v := Victim{}
	expectedLength := unsafe.Sizeof(v)

	b := &bytes.Buffer{}
	n, err := WriteObject(b, &v)
	assert.NoError(t, err)
	assert.EqualValues(t, expectedLength, n)
	assert.EqualValues(t, expectedLength, b.Len())
}

func TestByteBuilder(t *testing.T) {
	v := Victim{}
	expectedLength := unsafe.Sizeof(v)

	b := &ByteBuilder{}
	n, err := b.WriteObject(v)
	assert.NoError(t, err)
	assert.EqualValues(t, expectedLength, n)
	assert.EqualValues(t, expectedLength, b.Len())
}

func TestCastToByteArrayPointer(T *testing.T) {
	// https://stackoverflow.com/a/75018218
	v := any(Victim{})
	b := (*[unsafe.Sizeof(v)]byte)(unsafe.Pointer(&v))
	s := b[:]
	// test for inconsistency
	assert.NotEqualValues(T, int(unsafe.Sizeof(Victim{})), len(s))
}

func TestBinaryWrite(T *testing.T) {
	// https://stackoverflow.com/a/56063783
	v := any(Victim{})
	b := &bytes.Buffer{}
	err := binary.Write(b, binary.LittleEndian, v)
	assert.NoError(T, err)
	// test for inconsistency
	assert.NotEqualValues(T, int(unsafe.Sizeof(Victim{})), b.Len())
}
