package main

import (
	"hohice/go-webassembly/frontend/method"
	"syscall/js"
)

// Build with Go WASM fork
//go:generate rm -f ./assets/js/* ./assets/wasm/*
//go:generate bash -c "GOOS=js GOARCH=wasm go build -o ./assets/wasm/frontend.wasm frontend.go"
//go:generate bash -c "cp $DOLLAR(go env GOROOT)/misc/wasm/wasm_exec.js ./assets/js/wasm_exec.js"

// Integrate generated JS into a Go file for static loading.
//go:generate bash -c "go run assets_generate.go"

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(method.Add))
	js.Global().Set("subtract", js.NewCallback(method.Subtract))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
