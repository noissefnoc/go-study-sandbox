package cmd

import (
	"bytes"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"testing"
)

func TestShowNormal(t *testing.T) {
	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"show", "-i", "123", "-s", "日本語", "-b"}, want: "Integer option value: 123\nBoolean option value: true\n String option value: 日本語\n"},
	}

	for _, c := range testCases {
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := rwi.New(
			rwi.WithWriter(out),
			rwi.WithErrorWriter(errOut),
		)
		exit := Execute(ui, c.args)

		if exit != exitcode.Normal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Normal)
		}

		if out.String() != c.want {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), c.want)
		}

		if errOut.String() != "" {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), "")
		}
	}
}
