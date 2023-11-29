//go:build go1.18 && !go1.20

package bytebuilder

import (
	"reflect"
	"unsafe"
)

// SliceCast converts between different types of slices.
// Note: edge cases are not tested for, use with care.
func SliceCast[Ts, Td any](s []Ts) (d []Td) {
	// for testing size
	var sliceSrcMember Ts
	var sliceDestMember Td

	sliceSrcHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sliceDstHeader := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	sliceDstHeader.Cap = sliceSrcHeader.Cap * int(unsafe.Sizeof(sliceSrcMember)) / int(unsafe.Sizeof(sliceDestMember))
	sliceDstHeader.Len = sliceSrcHeader.Len * int(unsafe.Sizeof(sliceSrcMember)) / int(unsafe.Sizeof(sliceDestMember))
	sliceDstHeader.Data = sliceSrcHeader.Data

	return
}

// CarCdr returns a (readonly) object converted from your input slice, and the remaining (readonly) slice.
func CarCdr[T any](s []byte) (car *T, cdr []byte) {
	// for testing size
	var obj T
	objSize := unsafe.Sizeof(obj)

	srcSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	dstSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cdr))

	car = (*T)(unsafe.Pointer(srcSliceHeader.Data))

	dstSliceHeader.Cap = srcSliceHeader.Cap - int(objSize)
	dstSliceHeader.Len = srcSliceHeader.Len - int(objSize)
	dstSliceHeader.Data = srcSliceHeader.Data + objSize

	return
}
