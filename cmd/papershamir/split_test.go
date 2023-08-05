package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/bradenrayhorn/papershamir/internal/testutils"
	"github.com/matryer/is"
)

func TestSplit(t *testing.T) {
	is := is.New(t)

	out := bytes.NewBufferString("")
	cmd := SplitCmd{Parts: 2, Threshold: 2}

	err := cmd.Run(&CmdContext{in: bytes.NewBufferString("A"), out: out})
	is.NoErr(err)

	output := out.String()
	shares := strings.Split(output, "\n\n")
	is.Equal(len(shares), 2)

	for _, share := range shares {
		is.True(testutils.CreateShareRegexp(1, 2, 2).MatchString(share))
	}
}
