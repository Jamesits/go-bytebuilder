package bytebuilder

import (
	"reflect"
	"unsafe"
)

// emptyInterface is the header for an interface{} value.
//
// https://github.com/golang/go/blob/035db07d7c5f1b90ebc9bae03cab694685acebb8/src/reflect/value.go#L195-L199
// https://stackoverflow.com/a/57698257
type emptyInterface struct {
	typ, val unsafe.Pointer
}

// newArbitraryByteArray creates a (readonly) array from any memory location and length.
func newArbitraryByteArray(length uintptr, ptr uintptr) *[]byte {
	var ret []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = ptr
	return &ret
}
