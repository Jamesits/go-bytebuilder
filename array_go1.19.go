//go:build !go1.20

package bytebuilder

import (
	"reflect"
	"unsafe"
)

// NewArbitraryByteArray creates a (readonly) array from any memory location and length.
func NewArbitraryByteArray(length uintptr, ptr uintptr) *[]byte {
	var ret []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = ptr
	return &ret
}
