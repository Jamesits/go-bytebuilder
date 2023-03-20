package bytebuilder

import (
	"reflect"
	"unsafe"
)

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
