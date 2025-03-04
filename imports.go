//go:build !purego && !appengine && !wasm && !tinygo.wasm && !wasi
// +build !purego,!appengine,!wasm,!tinygo.wasm,!wasi

package wapc

func guestRequest(operationPtr uintptr, payloadPtr uintptr) {}

func guestResponse(ptr uintptr, len uint32) {}

func guestError(ptr uintptr, len uint32) {}

func hostCall(
	_ uintptr, _ uint32,
	_ uintptr, _ uint32,
	_ uintptr, _ uint32,
	_ uintptr, _ uint32,
) bool {
	return true
}

func hostResponseLen() uint32 { return 0 }

func hostResponse(ptr uintptr) {}

func hostErrorLen() uint32 { return 0 }

func hostError(ptr uintptr) {}

func consoleLog(ptr uintptr, size uint32) {}
