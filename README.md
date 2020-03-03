# waPC Guest Library for TinyGo

*Note:* Consider this SDK experimental.  We have yet to put it through more advanced use cases than "hello world".

This is the TinyGo implementation of the **waPC** standard for WebAssembly guest modules. It allows any waPC-compliant WebAssembly host to invoke to procedures inside a TinyGo compiled guest and similarly for the guest to invoke procedures exposed by the host.

## Example
The following is a simple example of synchronous, bi-directional procedure calls between a WebAssembly host runtime and the guest module.

```go
package main

import (
	wapc "github.com/wapc/wapc-guest-tinygo"
)

func main() {
	wapc.Register(wapc.Functions{
		"hello": hello,
	})
}

func hello(payload []byte) ([]byte, error) {
	wapc.HostCall("sample", "hello", []byte("Simon"))
	return []byte("Hello"), nil
}
```

```sh
tinygo build -o example/hello.wasm -target wasm -no-debug example/hello.go
```

## Known limitations

* Only go up to 1.13 is supported by TinyGo
* The `fmt` package requires `syscall/js.*` which are not imported by the waPC host
* TinyGo has limited `reflect` package support, thus libraries like protobuf will likely not work
