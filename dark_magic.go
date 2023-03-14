package bytebuilder

import (
	"reflect"
	"unsafe"
)

// newArbitraryByteArray creates a (readonly) array from any memory location and length.
func newArbitraryByteArray(length uintptr, ptr uintptr) *[]byte {
	var ret []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = ptr
	return &ret
}
