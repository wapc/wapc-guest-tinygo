package wapc

import (
	"reflect"
	"unsafe"
)

//go:wasm-module wapc
//go:export __guest_request
func guestRequest(operationPtr uintptr, payloadPtr uintptr)

//go:wasm-module wapc
//go:export __guest_response
func guestResponse(ptr uintptr, len uint32)

//go:wasm-module wapc
//go:export __guest_error
func guestError(ptr uintptr, len uint32)

//go:wasm-module wapc
//go:export __host_call
func hostCall(
	namespacePtr uintptr, namespaceLen uint32,
	operationPtr uintptr, operationLen uint32,
	payloadPtr uintptr, payloadLen uint32) bool

//go:wasm-module wapc
//go:export __host_response_len
func hostResponseLen() uint32

//go:wasm-module wapc
//go:export __host_response
func hostResponse(ptr uintptr)

//go:wasm-module wapc
//go:export __host_error_len
func hostErrorLen() uint32

//go:wasm-module wapc
//go:export __host_error
func hostError(ptr uintptr)

type (
	// Function is the function to register in your waPC module.
	Function func(payload []byte) ([]byte, error)

	// Functions is a map of function name to `Function`.
	Functions map[string]Function

	// HostError indicates an error when invoking a host operation.
	HostError struct {
		message string
	}
)

var (
	allFunctions = Functions{}
)

// Register adds functions by name to the registery.
// This should be invoked in `main()`.
func Register(functions Functions) {
	for name, fn := range functions {
		allFunctions[name] = fn
	}
}

//go:export __guest_call
func guestCall(operationSize uint32, payloadSize uint32) bool {
	operation := make([]byte, operationSize)
	payload := make([]byte, payloadSize)
	guestRequest(bytesToPointer(operation), bytesToPointer(payload))

	if f, ok := allFunctions[string(operation)]; ok {
		response, err := f(payload)
		if err != nil {
			message := err.Error()
			guestError(stringToPointer(message), uint32(len(message)))

			return false
		}

		guestResponse(bytesToPointer(response), uint32(len(response)))

		return true
	}

	message := `Could not find function "` + string(operation) + `"`
	guestError(stringToPointer(message), uint32(len(message)))

	return false
}

// HostCall invokes an operation on the host.  The host uses `namespace` and `operation`
// to route to the `payload` to the appropriate operation.  The host will return
// a response payload if successful.
func HostCall(namespace, operation string, payload []byte) ([]byte, error) {
	result := hostCall(
		stringToPointer(namespace), uint32(len(namespace)),
		stringToPointer(operation), uint32(len(operation)),
		bytesToPointer(payload), uint32(len(payload)),
	)
	if !result {
		errorLen := hostErrorLen()
		message := make([]byte, errorLen)
		hostError(bytesToPointer(message))

		return nil, &HostError{message: string(message)}
	}

	responseLen := hostResponseLen()
	response := make([]byte, responseLen)
	hostResponse(bytesToPointer(response))

	return response, nil
}

//go:inline
func bytesToPointer(s []byte) uintptr {
	return (*(*reflect.SliceHeader)(unsafe.Pointer(&s))).Data
}

//go:inline
func stringToPointer(s string) uintptr {
	return (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data
}

func (e *HostError) Error() string {
	return "Host error: " + e.message
}
