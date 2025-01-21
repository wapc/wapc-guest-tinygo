build:
	tinygo build -o example/hello.wasm -scheduler=none --no-debug -target=wasip1 -buildmode=c-shared example/hello.go
	$(MAKE) -C internal/e2e build

tests:
	go test -v ./...
	$(MAKE) -C internal test
