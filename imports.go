//go:build !purego && !appengine && !wasm && !tinygo.wasm && !wasi
// +build !purego,!appengine,!wasm,!tinygo.wasm,!wasi

package wapc

func guestRequest(operationPtr uintptr, payloadPtr uintptr) {}

func guestResponse(ptr uintptr, len uint32) {}

func guestError(ptr uintptr, len uint32) {}

func hostCall(
	bindingPtr uintptr, bindingLen uint32,
	namespacePtr uintptr, namespaceLen uint32,
	operationPtr uintptr, operationLen uint32,
	payloadPtr uintptr, payloadLen uint32) bool {
	return true
}

func hostResponseLen() uint32 { return 0 }

func hostResponse(ptr uintptr) {}

func hostErrorLen() uint32 { return 0 }

func hostError(ptr uintptr) {}

func consoleLog(ptr uintptr, size uint32) {}
