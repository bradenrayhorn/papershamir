package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/bradenrayhorn/papershamir/papershamir"
)

type CombineCmd struct{}

func (r *CombineCmd) Run(ctx *CmdContext) error {
	input, err := io.ReadAll(ctx.in)
	if err != nil {
		return err
	}
	shares := bytes.Split(bytes.ReplaceAll(bytes.TrimSpace(input), []byte("\r"), []byte{}), []byte("\n\n"))

	result, err := papershamir.Combine(shares, "")
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintf(ctx.out, "%s", result); err != nil {
		return err
	}

	return nil
}
