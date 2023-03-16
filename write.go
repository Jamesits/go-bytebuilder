//go:build go1.18

package bytebuilder

import (
	"io"
	"unsafe"
)

// WriteObject write any object's internal memory representation into an io.Writer.
// Arguments:
// - w: io.Writer
// - object: a pointer to the source object
func WriteObject[T any](w io.Writer, object *T) (n int, err error) {
	objectSize := unsafe.Sizeof(*object)
	s := newArbitraryByteArray(objectSize, uintptr(unsafe.Pointer(object)))
	return w.Write(*s)
}
