GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

wasmjs:
	GOOS=js GOARCH=wasm go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o static/wasmjs/tags.wasm cmd/wasmjs/main.go
