build:
	tinygo build -o example/hello.wasm -scheduler=none --no-debug -target=wasi example/hello.go
	$(MAKE) -C internal/e2e build

tests:
	go test -v ./...
	$(MAKE) -C internal test
