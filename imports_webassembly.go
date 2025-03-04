//go:build wasm || tinygo.wasm || wasi || wasip1
// +build wasm tinygo.wasm wasi wasip1

package wapc

//go:wasmimport wapc __guest_request
func guestRequest(operationPtr uintptr, payloadPtr uintptr)

//go:wasmimport wapc __guest_response
func guestResponse(ptr uintptr, len uint32)

//go:wasmimport wapc __guest_error
func guestError(ptr uintptr, len uint32)

//go:wasmimport wapc __host_call
func hostCall(
	bindingPtr uintptr, bindingLen uint32,
	namespacePtr uintptr, namespaceLen uint32,
	operationPtr uintptr, operationLen uint32,
	payloadPtr uintptr, payloadLen uint32) bool

//go:wasmimport wapc __host_response_len
func hostResponseLen() uint32

//go:wasmimport wapc __host_response
func hostResponse(ptr uintptr)

//go:wasmimport wapc __host_error_len
func hostErrorLen() uint32

//go:wasmimport wapc __host_error
func hostError(ptr uintptr)

//go:wasmimport wapc __console_log
func consoleLog(ptr uintptr, size uint32)
