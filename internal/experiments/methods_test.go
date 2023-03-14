package experiments

import (
	"bytes"
	"encoding/binary"
	"github.com/jamesits/go-bytebuilder"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

type Victim struct {
	A bool
	B uint64
	C bool
}

func testWrapper(T *testing.T, actualSize int) {
	expectedSize := int(reflect.TypeOf(Victim{}).Size())
	assert.EqualValues(T, expectedSize, actualSize)
}

func TestByteBuilder(T *testing.T) {
	// https://stackoverflow.com/a/57698257
	v := any(Victim{})
	b := bytebuilder.ByteBuilder{}
	_, _ = b.WriteObject(v)
	testWrapper(T, b.Len())
}

// should fail
func TestCastToByteArrayPointer(T *testing.T) {
	// https://stackoverflow.com/a/75018218
	v := any(Victim{})
	b := (*[unsafe.Sizeof(v)]byte)(unsafe.Pointer(&v))
	s := b[:]
	testWrapper(T, len(s))
}

// should fail
func TestBinaryWrite(T *testing.T) {
	// https://stackoverflow.com/a/56063783
	v := any(Victim{})
	b := &bytes.Buffer{}
	_ = binary.Write(b, binary.LittleEndian, v)
	testWrapper(T, b.Len())
}
