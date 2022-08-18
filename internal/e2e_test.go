package internal_test

import (
	"context"
	_ "embed"
	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/wazero/api"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

// testCtx is an arbitrary, non-default context. Non-nil also prevents linter errors.
var testCtx = context.WithValue(context.Background(), struct{}{}, "arbitrary")

// consoleLogWasm was compiled from testdata/__console_log/main.go
//
//go:embed testdata/__console_log/main.wasm
var consoleLogWasm []byte

func Test_EndToEnd(t *testing.T) {
	type testCase struct {
		name  string
		guest []byte
		test  func(t *testing.T, guest api.Module, host *wapcHost)
	}

	tests := []testCase{
		{
			name:  "ConsoleLog",
			guest: consoleLogWasm,
			test: func(t *testing.T, guest api.Module, host *wapcHost) {
				// main invokes ConsoleLog
				require.Equal(t, []string{"msg", "msg1", "msg"}, host.consoleLogMessages)
			},
		},
	}

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfig().
		// WebAssembly 2.0 allows use of any version of TinyGo, including 0.24+.
		WithWasmCore2())
	defer r.Close(testCtx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements system I/O such as console output and
	// is required for `tinygo build -target=wasi`
	if _, err := wasi_snapshot_preview1.Instantiate(testCtx, r); err != nil {
		t.Errorf("Error instantiating WASI - %v", err)
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			h, host := instantiateWapcHost(t, r)
			defer host.Close(testCtx)

			g, err := r.InstantiateModuleFromBinary(testCtx, tc.guest)
			if err != nil {
				t.Errorf("Error instantiating waPC guest - %v", err)
			}
			defer g.Close(testCtx)

			tc.test(t, g, h)
		})
	}
}
