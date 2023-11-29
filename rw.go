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
	s := NewArbitraryByteArray(objectSize, uintptr(unsafe.Pointer(object)))
	return w.Write(*s)
}

// ReadPartial reads from an io.Reader directly into an object.
func ReadPartial[T any](r io.Reader, object *T) (n int, err error) {
	objectSize := unsafe.Sizeof(*object)
	s := NewArbitraryByteArray(objectSize, uintptr(unsafe.Pointer(object)))
	return r.Read(*s)
}

// Skip certain length from the reader.
func Skip(reader io.Reader, count int64) (int64, error) {
	// https://stackoverflow.com/a/20330589
	switch r := reader.(type) {
	case io.Seeker:
		return r.Seek(count, io.SeekCurrent)
	default:
		return io.CopyN(io.Discard, r, count)
	}
}

// Copy an object to a byte buffer.
func Copy[T any](b []byte, object *T) (n int, err error) {
	objectSize := unsafe.Sizeof(*object)
	s := NewArbitraryByteArray(objectSize, uintptr(unsafe.Pointer(object)))
	return copy(b, *s), nil
}
