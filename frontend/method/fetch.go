// +build js,wasm

package method

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"syscall/js"
)

type writer js.Value

// Write implements io.Writer.
func (d writer) Write(p []byte) (n int, err error) {
	node := document.Call("createElement", "div")
	node.Set("textContent", string(p))
	js.Value(d).Call("appendChild", node)
	return len(p), nil
}

func Fetch(i []js.Value) {
	t := document.Call("getElementById", "fetch")
	logger := log.New((*writer)(&t), "", log.LstdFlags)

	c := http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://httpbin.org/anything",
		strings.NewReader(`{"test":"test"}`),
	)
	if err != nil {
		logger.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		logger.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Print(string(b))
}
