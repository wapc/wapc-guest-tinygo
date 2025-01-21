# waPC Guest Library for TinyGo

This is the TinyGo implementation of the **waPC** standard for WebAssembly guest modules. It allows any waPC-compliant WebAssembly host to invoke to procedures inside a TinyGo compiled guest and similarly for the guest to invoke procedures exposed by the host.

## Example
The following is a simple example of synchronous, bi-directional procedure calls between a WebAssembly host runtime and the guest module.

> It is recommended to use the latest versions Go and TinyGo

For TinyGo 0.35+

```go
package main

import (
	wapc "github.com/wapc/wapc-guest-tinygo"
)

//go:wasmexport wapc_init
func Initialize() {
	wapc.RegisterFunctions(wapc.Functions{
		"hello": hello,
	})
}

func hello(payload []byte) ([]byte, error) {
	wapc.HostCall("myBinding", "sample", "hello", []byte("Simon"))
	return []byte("Hello"), nil
}
```

```sh
tinygo build -o example/hello.wasm -scheduler=none --no-debug -target=wasip1 -buildmode=c-shared example/hello.go
```
