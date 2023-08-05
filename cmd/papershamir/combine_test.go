package main

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestCombine(t *testing.T) {
	is := is.New(t)

	in := bytes.NewBufferString("99 4E XNKN413A\n      XNKN413A\n\n97 K7 E3A87326\n      E3A87326")
	out := bytes.NewBufferString("")
	cmd := CombineCmd{}
	err := cmd.Run(&CmdContext{in, out})
	is.NoErr(err)

	output := out.String()
	is.Equal(output, "x")
}
