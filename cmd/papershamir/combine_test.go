package main

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestCombine(t *testing.T) {
	is := is.New(t)

	in := bytes.NewBufferString("994E XNKN 413A\nXNKN 413A\n\n97K7 E3A8 7326\nE3A8 7326")
	out := bytes.NewBufferString("")
	cmd := CombineCmd{}
	err := cmd.Run(&CmdContext{in, out})
	is.NoErr(err)

	output := out.String()
	is.Equal(output, "x")
}
