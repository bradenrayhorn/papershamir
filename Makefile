.PHONY: wasm

SHELL=/bin/bash

wasm:
	GOOS=js GOARCH=wasm go build -o ./web/public/main.wasm ./cmd/web
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./web/public/
