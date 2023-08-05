package main

import (
	"io"
	"os"

	"github.com/alecthomas/kong"
)

type CmdContext struct {
	in  io.Reader
	out io.Writer
}

var cli struct {
	Split   SplitCmd   `cmd:"" help:"Split secret into parts. Pass secret as stdin."`
	Combine CombineCmd `cmd:"" help:"Combine parts into secret. Pass parts as stdin."`
}

func main() {
	ctx := kong.Parse(&cli, kong.Bind(os.Stdin))
	err := ctx.Run(&CmdContext{in: os.Stdin, out: os.Stdout})
	ctx.FatalIfErrorf(err)
}
