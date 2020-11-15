.PHONY: all build run clean

all: build run

static/mywasm.wasm: wasm/mywasm.go
	cd wasm && GOOS=js GOARCH=wasm go build -o ../static/mywasm.wasm

server: server.go
	go build server.go

build: static/mywasm.wasm server
	@echo "Built all"

run: build
	@echo "Run server..."
	./server

clean:
	@echo "Clean up..."
	rm -f static/mywasm.wasm
	rm -f server
