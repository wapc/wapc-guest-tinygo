package wapc

import (
	"reflect"
	"unsafe"
)

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

// RegisterFunctions adds functions by name to the registry.
// This should be invoked in `main()`.
func RegisterFunctions(functions Functions) {
	for name, fn := range functions {
		allFunctions[name] = fn
	}
}

// RegisterFunction adds a single function by name to the registry.
// This should be invoked in `main()`.
func RegisterFunction(name string, fn Function) {
	allFunctions[name] = fn
}

//go:export __guest_call
func guestCall(operationSize uint32, payloadSize uint32) bool {
	operation := make([]byte, operationSize) // alloc
	payload := make([]byte, payloadSize)     // alloc
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

// ConsoleLog writes the message the underlying waPC console logger.
func ConsoleLog(msg string) {
	consoleLog(stringToPointer(msg), uint32(len(msg)))
}

// HostCall invokes an operation on the host.  The host uses `namespace` and `operation`
// to route to the `payload` to the appropriate operation.  The host will return
// a response payload if successful.
func HostCall(binding, namespace, operation string, payload []byte) ([]byte, error) {
	result := hostCall(
		stringToPointer(binding), uint32(len(binding)),
		stringToPointer(namespace), uint32(len(namespace)),
		stringToPointer(operation), uint32(len(operation)),
		bytesToPointer(payload), uint32(len(payload)),
	)
	if !result {
		errorLen := hostErrorLen()
		message := make([]byte, errorLen) // alloc
		hostError(bytesToPointer(message))

		return nil, &HostError{message: string(message)} // alloc
	}

	responseLen := hostResponseLen()
	response := make([]byte, responseLen) // alloc
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
	return "Host error: " + e.message // alloc
}
