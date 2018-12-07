include .env

generate:
	go generate -x ./frontend/
wasm:
	GOARCH=wasm GOOS=js go build -o frontend.wasm frontend/frontend.go
serve:
	CGO_ENABLED=0 go run server.go
clean:
	go clean
	@rm -f frontend/assets/js/* frontend/assets/wasm/* frontend/bundle/*

.PHONY: wasm run