package set_test

import (
	"testing"

	"github.com/goware/set"
)

func TestInt64Set(t *testing.T) {
	s := set.NewInt64Set(1, 2, 3, 3)
	s = s.Add(5, 5, 3, 2, 1, 4, 4, 5)

	if len(s) != 5 {
		t.Error("expecting 5 values in the set")
	}
}

func TestStringSet(t *testing.T) {
	s := set.NewStringSet("a", "b", "a", "c")
	s = s.Add("d", "a", "b", "c", "e")

	if len(s) != 5 {
		t.Error("expecting 5 values in the set")
	}
}
