package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/bradenrayhorn/papershamir/papershamir"
)

type SplitCmd struct {
	Parts     int `help:"Number of parts to split secret into." default:"5"`
	Threshold int `help:"Number of parts required to reassemble the secret." default:"3"`
}

func (r *SplitCmd) Run(ctx *CmdContext) error {
	secret, err := io.ReadAll(ctx.in)
	if err != nil {
		return err
	}
	secret = bytes.TrimSpace(secret)

	shares, err := papershamir.Split(secret, r.Parts, r.Threshold, "test key")
	if err != nil {
		return err
	}

	for i, share := range shares {
		if _, err := fmt.Fprintf(ctx.out, "%s", share); err != nil {
			return err
		}

		if i < len(shares)-1 {
			if _, err := fmt.Fprintln(ctx.out); err != nil {
				return err
			}
		}
	}

	return nil
}
