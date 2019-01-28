package main

import (
	"github.com/noissefnoc/go-study-sandbox/cli/cobra/cmd"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"os"
)

func run(ui *rwi.RWI)exitcode.ExitCode {
	ui.Output("Hello world")
	return exitcode.Normal
}

func main() {
	cmd.Execute(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(os.Stdout),
			rwi.WithErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Exit()
}
