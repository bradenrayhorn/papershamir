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
	js.Global().Set("shamirCombineQR", js.FuncOf(shamirCombineQR))

	<-c
}

func shamirSplit(this js.Value, args []js.Value) any {
	passphrase, err := papershamir.RandomPassphrase()
	if err != nil {
		return makeJsError(err)
	}

	secret := args[0].String()
	parts := args[1].Int()
	threshold := args[2].Int()

	hexrShares, err := papershamir.SplitHexr([]byte(secret), parts, threshold, sanitizePassphrase(passphrase))
	if err != nil {
		return makeJsError(err)
	}

	textResult := make([]any, len(hexrShares), len(hexrShares))
	for i := range hexrShares {
		textResult[i] = hexrShares[i]
	}

	qrShares, err := papershamir.SplitQR([]byte(secret), parts, threshold, sanitizePassphrase(passphrase))
	if err != nil {
		return makeJsError(err)
	}

	qrResult := make([]any, len(qrShares), len(qrShares))
	for i := range qrResult {
		qrResult[i] = qrShares[i]
	}

	// assemble response
	obj := map[string]any{}
	obj["text"] = textResult
	obj["qr"] = qrResult
	obj["passphrase"] = passphrase

	return js.ValueOf(obj)
}

func shamirCombine(this js.Value, args []js.Value) any {
	input := []byte(args[0].String())
	key := args[1].String()
	shares := bytes.Split(bytes.ReplaceAll(bytes.TrimSpace(input), []byte("\r"), []byte{}), []byte("\n\n"))
	secret, err := papershamir.CombineHexr(shares, sanitizePassphrase(key))
	if err != nil {
		return js.Global().Get("Error").New(err.Error())
	}

	return string(secret)
}

func shamirCombineQR(this js.Value, args []js.Value) any {
	input := args[0]
	key := args[1].String()
	shares := make([][]byte, 0)
	for i := 0; i < input.Length(); i++ {
		shares = append(shares, []byte(input.Index(i).String()))
	}
	secret, err := papershamir.CombineQR(shares, sanitizePassphrase(key))
	if err != nil {
		return js.Global().Get("Error").New(err.Error())
	}

	return string(secret)
}

func makeJsError(err error) js.Value {
	return js.Global().Get("Error").New(err.Error())
}

func sanitizePassphrase(passphrase string) string {
	return strings.ToUpper(
		strings.TrimSpace(
			strings.ReplaceAll(passphrase, " ", ""),
		),
	)
}
