tinygo_sources := $(wildcard testdata/*/*.go)
.PHONY: build.wasm
build-wasm: $(tinygo_sources)
	@echo "------------------"
	@echo "Building Test Wasm"
	@echo "------------------"
	@for f in $^; do \
	    tinygo build -o $$(echo $$f | sed -e 's/\.go/\.wasm/') -scheduler=none --no-debug -target=wasi $$f; \
	done
