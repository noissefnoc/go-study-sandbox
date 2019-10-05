package main

import (
	"github.com/alecthomas/kong"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/repr"
	"strings"
)

var digitLexer = lexer.Must(lexer.Regexp(
	`(?m)` +
		`(?P<DigitExcludeZero>)[1-9]` +
		`|(?P<Zero>)0`,
))

type Number struct {
	Value float64 `@DigitExcludeZero { @(Zero|DigitExcludeZero) }`
}

func main() {
	var cli struct {
		Expr []string `arg required help:"Expression to parse."`
	}
	ctx := kong.Parse(&cli)
	p, err := participle.Build(&Number{},
		participle.Lexer(digitLexer))
	ctx.FatalIfErrorf(err)

	expr := &Number{}
	err = p.ParseString(strings.Join(cli.Expr, " "), expr)
	ctx.FatalIfErrorf(err)

	repr.Println(expr)
}
