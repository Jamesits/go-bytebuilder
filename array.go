//go:build go1.20

package bytebuilder

import "unsafe"

// NewArbitraryByteArray creates a (readonly) array from any memory location and length.
// It assumes the memory behind the ptr is allocated elsewhere, readonly and pinned.
//
//go:uintptrescapes
func NewArbitraryByteArray(length uintptr, ptr uintptr) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(ptr)), length)
}
