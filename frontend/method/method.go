// +build js,wasm

package method

import (
	"strconv"
	"syscall/js"
)

var document js.Value

func init() {
	document = js.Global().Get("document")
}

func Add(i []js.Value) {
	value1 := document.Call("getElementById", i[0].String()).Get("value").String()
	value2 := document.Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	document.Call("getElementById", i[2].String()).Set("value", int1+int2)
}

func Subtract(i []js.Value) {
	value1 := document.Call("getElementById", i[0].String()).Get("value").String()
	value2 := document.Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	document.Call("getElementById", i[2].String()).Set("value", int1-int2)
}

func sum(s []int, c chan<- int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum
}

//Sum to modify <div id="sum"/> and call createElement
func Sum( /*i []js.Value*/ ) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	div := document.Call("getElementById", "sum")

	node := document.Call("createElement", "div")
	node.Set("innerText", x+y)

	div.Call("appendChild", node)

}
