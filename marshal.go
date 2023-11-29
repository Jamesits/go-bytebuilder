//go:build go1.18

package bytebuilder

import (
	"bytes"
	"errors"
	"unsafe"
)

var SizeMismatch = errors.New("size mismatch")

// Marshal returns the exact internal memory representation of v.
func Marshal[T any](v *T) ([]byte, error) {
	var b bytes.Buffer
	_, err := WriteObject(&b, v)
	return b.Bytes(), err
}

// Unmarshal puts the data into the value pointed to by v.
func Unmarshal[T any](data []byte, v *T) (err error) {
	dataSize := uintptr(len(data))
	objectSize := unsafe.Sizeof(*v)
	if dataSize != objectSize {
		err = SizeMismatch
	}

	s := NewArbitraryByteArray(objectSize, uintptr(unsafe.Pointer(v)))
	copy(*s, data) // copy() does size check itself
	return
}
