//go:build js && wasm

package main

import (
	"bytes"
	"strings"
	"syscall/js"

	"github.com/bradenrayhorn/papershamir/papershamir"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("shamirSplit", js.FuncOf(shamirSplit))
	js.Global().Set("shamirCombine", js.FuncOf(shamirCombine))

	<-c
}

func shamirSplit(this js.Value, args []js.Value) any {
	secret := args[0].String()
	parts := args[1].Int()
	threshold := args[2].Int()
	shares, err := papershamir.Split([]byte(secret), parts, threshold)
	if err != nil {
		return js.Global().Get("Error").New(err.Error())
	}

	return strings.Join(shares, "\n")
}

func shamirCombine(this js.Value, args []js.Value) any {
	input := []byte(args[0].String())
	shares := bytes.Split(bytes.ReplaceAll(bytes.TrimSpace(input), []byte("\r"), []byte{}), []byte("\n\n"))
	secret, err := papershamir.Combine(shares)
	if err != nil {
		return js.Global().Get("Error").New(err.Error())
	}

	return string(secret)
}
