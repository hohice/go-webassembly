package main

import (
	"flag"
	"hohice/go-webassembly/frontend/bundle"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":9090", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(bundle.Assets)))
}
