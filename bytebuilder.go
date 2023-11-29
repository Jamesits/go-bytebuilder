package bytebuilder

import (
	"bytes"
	"reflect"
	"unsafe"
)

// ByteBuilder is a `bytes.Buffer` with the capability of writing any object into it.
type ByteBuilder struct {
	bytes.Buffer
}

// WriteObject writes any object's internal memory representation into the buffer.
// Arguments:
// - obj: the source object itself
func (b *ByteBuilder) WriteObject(obj any) (n int, err error) {
	objectSize := reflect.TypeOf(obj).Size()
	s := NewArbitraryByteArray(objectSize, uintptr((*emptyInterface)(unsafe.Pointer(&obj)).val))
	return b.Write(s)
}
