package bytebuilder

import "unsafe"

// emptyInterface is the header for an interface{} value.
// https://github.com/golang/go/blob/035db07d7c5f1b90ebc9bae03cab694685acebb8/src/reflect/value.go#L195-L199
type emptyInterface struct {
	typ, val unsafe.Pointer
}
