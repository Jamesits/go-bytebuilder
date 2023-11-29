//go:build go1.20

package bytebuilder

import (
	"unsafe"
)

// SliceCast converts between different types of slices.
// It assumes the source slice is created from a readonly pinned memory area.
// Note: edge cases are not tested for, use with care.
func SliceCast[Ts, Td any](s []Ts) (d []Td) {
	// for testing size
	var sliceSrcMember Ts
	var sliceDestMember Td

	return unsafe.Slice((*Td)(unsafe.Pointer(&s[0])), len(s)*int(unsafe.Sizeof(sliceSrcMember))/int(unsafe.Sizeof(sliceDestMember)))
}

// CarCdr returns a (readonly) object converted from your input slice, and the remaining (readonly) slice.
// It assumes the slice is created from a readonly pinned memory area.
func CarCdr[T any](s []byte) (car *T, cdr []byte) {
	// for testing size
	var obj T
	objSize := unsafe.Sizeof(obj)

	return (*T)(unsafe.Pointer(&s[0])), unsafe.Slice(&s[objSize], len(s)-int(objSize))
}
