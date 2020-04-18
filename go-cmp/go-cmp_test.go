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

func checkResult(t *testing.T, expected, actual interface{}) {
	t.Helper()

	ignoreField := cmpopts.IgnoreFields(Person{}, "Name")

	if diff := cmp.Diff(expected, actual, ignoreField); diff != "" {
		t.Errorf("Mismatch (-want +got)\n%s", diff)
	}
}
