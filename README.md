# waPC Guest Library for TinyGo

This is the TinyGo implementation of the **waPC** standard for WebAssembly guest modules. It allows any waPC-compliant WebAssembly host to invoke to procedures inside a TinyGo compiled guest and similarly for the guest to invoke procedures exposed by the host.

## Example
The following is a simple example of synchronous, bi-directional procedure calls between a WebAssembly host runtime and the guest module.

```go
package main

import (
	wapc "github.com/wapc/wapc-guest-tinygo"
)

func main() {
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
tinygo build -o example/hello.wasm -scheduler=none --no-debug -target=wasi example/hello.go
```

## Considerations

* It is recommended to use the latest versions Go and TinyGo
* Avoid using the `fmt` package as it requires `syscall/js.*` which are not implemented by the waPC host libraries
* TinyGo has limited `reflect` package support, thus generated Protobuf code will likely not work without some tweaking (But we have gotten it to work!)
