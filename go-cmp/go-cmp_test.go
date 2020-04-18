package go_cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Person struct {
	Name string
	Age  int
}

type Hoge struct {
	Moji          string
	AnotherStruct *Huga
	Num           int
	Flag          bool
}

type Huga struct {
	Moji string
}

func TestGoCmp(t *testing.T) {
	p1 := &Person{
		Name: "Alice",
		Age:  10,
	}

	p2 := &Person{
		Name: "Jone",
		Age:  10,
	}

	checkResult(t, p1, p2)
}

func TestNestedStructure(t *testing.T) {
	h1 := &Hoge{
		Moji: "hoge",
		AnotherStruct: &Huga{
			Moji: "huga",
		},
		Num:  1,
		Flag: false,
	}

	h2 := &Hoge{
		Moji: "hogehoge",
		AnotherStruct: &Huga{
			Moji: "hugahoge",
		},
		Num:  2,
		Flag: false,
	}

	checkResult(t, h1, h2)
}

func checkResult(t *testing.T, expected, actual interface{}) {
	t.Helper()
	opts := buildCmpOptions(t)
	if diff := cmp.Diff(expected, actual, opts...); diff != "" {
		t.Errorf("Mismatch (-want +got)\n%s", diff)
	}
}

func buildCmpOptions(t *testing.T) []cmp.Option {
	t.Helper()
	opts := []cmp.Option{
		cmpopts.IgnoreFields(Person{}, "Name"),
		cmpopts.IgnoreFields(Huga{}, "Moji"),
		cmpopts.IgnoreFields(Hoge{}, "Num", "Moji"),
	}
	return opts
}
