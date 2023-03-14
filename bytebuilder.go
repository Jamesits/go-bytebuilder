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
func (b *ByteBuilder) WriteObject(src any) (n int, err error) {
	bLength := reflect.TypeOf(src).Size()
	var s []uint8
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = int(bLength)
	sh.Len = int(bLength)
	sh.Data = uintptr((*emptyInterface)(unsafe.Pointer(&src)).val)
	return b.Write(s)
}
