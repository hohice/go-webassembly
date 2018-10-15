include .env

generate:
	go generate -x ./frontend/
serve:
	CGO_ENABLED=0 go run server.go
clean:
	go clean
	@rm -f frontend/assets/js/* frontend/assets/wasm/* frontend/bundle/*

.PHONY: wasm run