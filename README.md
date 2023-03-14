# ByteBuilder

A `bytes.Buffer` with the capability of writing any object's exact memory layout (with every padding) into it.

[![Go Reference](https://pkg.go.dev/badge/github.com/Jamesits/go-bytebuilder.svg)](https://pkg.go.dev/github.com/Jamesits/go-bytebuilder)

## Usage

```go
package main

import (
	"fmt"
	"github.com/jamesits/go-bytebuilder"
	"unsafe"
)

type SomeStruct struct {
	Value1 uint16
	// there is an implicit padding
	Value2 uint64
}

func main() {
	s := SomeStruct{
		Value1: 100,
		Value2: 500,
	}
	fmt.Printf("size of s is %d\n", unsafe.Sizeof(s))
	
	// using the ByteBuilder
	var bb bytebuilder.ByteBuilder
	_, _ = bb.WriteObject(s)
	fmt.Printf("size of bb is %d\n", bb.Len())
	
	// using the Writer wrapper
	b := &bytes.Buffer{}
	_, _ = bytebuilder.WriteObject(b, &s)
	fmt.Printf("size of b is %d\n", b.Len())
}
```

Other methods of a `ByteBuilder` works exactly the same as a `bytes.Buffer`.
