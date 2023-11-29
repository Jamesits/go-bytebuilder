//go:build go1.20

package bytebuilder

import "unsafe"

// NewArbitraryByteArray creates a (readonly) array from any memory location and length.
func NewArbitraryByteArray(length uintptr, ptr uintptr) *[]byte {
	b := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), length)
	return &b
}
