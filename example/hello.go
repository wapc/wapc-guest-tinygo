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
