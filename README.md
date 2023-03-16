# ByteBuilder

A `bytes.Buffer` with the capability of writing any object's exact memory layout (with every padding) into it.

[![Go Reference](https://pkg.go.dev/badge/github.com/Jamesits/go-bytebuilder.svg)](https://pkg.go.dev/github.com/Jamesits/go-bytebuilder)

## WARNING

This is not a proper serialization protocol implementation, and has never intended to be one. Any data acquired with 
this method is usually only meaningful under the same CPU architecture, the same OS, AND the same compiler.

Cases when you should NOT use this library:
- If you want to dump structs to a persistent storage only to restore them later in the same program, use [encoding/gob](https://pkg.go.dev/encoding/gob).
- To encode/decode between basic types and a binary buffer, use [encoding/binary](https://pkg.go.dev/encoding/binary).
- For reading/writing machine-readable config files, use [mapstructure](https://pkg.go.dev/github.com/mitchellh/mapstructure), [encoding/json](https://pkg.go.dev/encoding/json), or [encoding/xml](https://pkg.go.dev/encoding/xml).
- For exchanging data between network services, use [Protocol Buffers](https://pkg.go.dev/google.golang.org/protobuf) or [Cap'n Proto](https://github.com/capnproto/go-capnp).

## Usage

### `bytebuilder.WriteObject`

Takes a `io.Writer` and a pointer to an object. Writes the object into the Writer. 

`bytebuilder.WriteObject` is well-defined (relies only on public methods) although unsafe (bypasses type safety). 
Generics support is required, so this function is only available on Golang 1.18+.

```go
package main

import (
	"bytes"
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
	
	// using the Writer wrapper
	var b bytes.Buffer
	_, _ = bytebuilder.WriteObject(&b, &s)
	fmt.Printf("size of b is %d\n", b.Len())
	
	// you can write more data into the same buffer
	b.Write([]byte{1, 1, 4, 5, 1, 4})
}
```

### `bytebuilder.ByteBuilder`

`ByteBuilder` provides a `StringBuilder`-like API to concentrate multiple objects into one big buffer. 

`bytebuilder.ByteBuilder{}.WriteObject` writes the object into the buffer. It relies on internal knowledge of the 
Golang runtime implementation and unsafe. It does not require generics, and works on all recent Golang versions. 

Other methods of a `ByteBuilder` works exactly the same as a `bytes.Buffer`, so you can use `Write()` or other methods 
for certain use cases.

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
	
	// you can write more objects into it
	_, _ = bb.WriteObject(s)
	
	// or use it as a generic buffer
	_, _ = bb.Write([]byte{1, 1, 4, 5, 1, 4})
}
```

### Real-world Examples

Say we have the following type definitions in C:

```c
typedef struct {
    uint64_t id;
    char name[32];
} student;

typedef struct {
    char name[32];
    uint8_t student_count;
    student students[1];
} class;
```

And we have a C function that takes an argument of type `*class`.

It is impossible to express structs with [flexible array members](https://en.wikipedia.org/wiki/Flexible_array_member) 
natively in Golang. But if your C compiler's padding rules happens to be the same as Golang's, this library offers you 
a shortcut besides your regular byte operations. 

```go
package main

import (
	"bytes"
	"github.com/jamesits/go-bytebuilder"
)

type Student struct {
	Id uint64
	Name [32]byte
}

type Class struct {
	Name [32]byte
	StudentCount uint8
}

func main() {
	class := Class {
		Name: [32]byte{'c', 'l', 'a', 's', 's'},
		StudentCount: 2,
    }
	
	alice := Student {
		Id: 1,
		Name: [32]byte{'A', 'l', 'i', 'c', 'e'},
    }
	
	bob := Student {
		Id: 2,
		Name: [32]byte{'B', 'o', 'b'},
    }
	
	var buf bytes.Buffer
	_, _ = bytebuilder.WriteObject(&buf, &class)
	_, _ = bytebuilder.WriteObject(&buf, &alice)
	_, _ = bytebuilder.WriteObject(&buf, &bob)
	
	// now buf.Bytes() contains exactly what you need to give to the native function
	fmt.Printf("%v\n", buf.Bytes())
}
```
