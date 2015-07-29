package set

import (
	"sort"

	"github.com/xtgo/set"
)

// Int64 set
type Int64Set []int64

func NewInt64Set(v ...int64) Int64Set {
	s := Int64Set{}
	if len(v) > 0 {
		s = s.Add(v...)
	}
	return s
}

func (s Int64Set) Len() int           { return len(s) }
func (s Int64Set) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int64Set) Less(i, j int) bool { return s[i] < s[j] }

func (s Int64Set) Union(t Int64Set) Int64Set {
	data := append(s, t...)
	n := set.Union(data, len(s))
	return data[:n]
}

func (s Int64Set) Add(v ...int64) Int64Set {
	data := append(s, v...)
	n := set.Union(data, len(s))
	return data[:n]
}

// String set
type StringSet []string

func NewStringSet(v ...string) StringSet {
	s := StringSet{}
	if len(v) > 0 {
		s = s.Add(v...)
	}
	return s
}

func (s StringSet) Union(t StringSet) StringSet {
	data := append(s, t...)
	n := set.Union(sort.StringSlice(data), len(s))
	return data[:n]
}

func (s StringSet) Add(v ...string) StringSet {
	data := append(s, v...)
	n := set.Union(sort.StringSlice(data), len(s))
	return data[:n]
}
