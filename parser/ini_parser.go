package main

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/repr"
	"os"
)

var iniLexer = lexer.Must(lexer.Regexp(
	`(?m)` +
		`(\s+)` +
		`|(^[#;].*$)` +
		`|(?P<Ident>[a-zA-Z][a-zA-Z_\d]*)` +
		`|(?P<String>"(?:\\.|[^"])*")` +
		`|(?P<Float>\d+(?:\.\d+)?)` +
		`|(?P<Punct>[][=])`,
))

type INI struct {
	Properties []*Property `@@*`
	Sections   []*Section  `@@*`
}

type Section struct {
	Identifier string      `"[" @Ident "]"`
	Properties []*Property `@@*`
}

type Property struct {
	Key   string `@Ident "="`
	Value *Value `@@`
}

type Value struct {
	String *string  `  @String`
	Number *float64 `| @Float`
}

func main() {
	parser, err := participle.Build(&INI{},
		participle.Lexer(iniLexer),
		participle.Unquote("String"),
	)

	if err != nil {
		fmt.Printf("EBNF build failed: %v\n", err)
		os.Exit(1)
	}

	ini := &INI{}
	err = parser.ParseString(`
age = 21
name = "Bob Smith"

[address]
city = "Beverly Hills"
postal_code = 90210
`, ini)

	if err != nil {
		fmt.Printf("parse failed: %v\n", err)
		os.Exit(1)
	}

	repr.Println(ini, repr.Indent("  "), repr.OmitEmpty(true))
}