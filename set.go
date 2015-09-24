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

func (s Int64Set) Do(op set.Op, t Int64Set) Int64Set {
	data := append(s, t...)
	n := op(data, len(s))
	ss := data[:n]
	sort.Sort(s)
	return ss
}

func (s Int64Set) Union(t Int64Set) Int64Set  { return s.Do(set.Union, t) }
func (s Int64Set) Inter(t Int64Set) Int64Set  { return s.Do(set.Inter, t) }
func (s Int64Set) Diff(t Int64Set) Int64Set   { return s.Do(set.Diff, t) }
func (s Int64Set) Add(v ...int64) Int64Set    { return s.Union(Int64Set(v)) }
func (s Int64Set) Remove(v ...int64) Int64Set { return s.Diff(NewInt64Set(v...)) }
func (s Int64Set) Exists(v int64) bool        { return len(s.Inter(Int64Set{v})) != 0 }

// String set
type StringSet []string

func NewStringSet(v ...string) StringSet {
	s := StringSet{}
	if len(v) > 0 {
		s = s.Add(v...)
	}
	return s
}

func (s StringSet) Do(op set.Op, t StringSet) StringSet {
	data := append(s, t...)
	n := op(sort.StringSlice(data), len(s))
	return data[:n]
}

func (s StringSet) Union(t StringSet) StringSet  { return s.Do(set.Union, t) }
func (s StringSet) Inter(t StringSet) StringSet  { return s.Do(set.Inter, t) }
func (s StringSet) Diff(t StringSet) StringSet   { return s.Do(set.Diff, t) }
func (s StringSet) Add(v ...string) StringSet    { return s.Union(StringSet(v)) }
func (s StringSet) Remove(v ...string) StringSet { return s.Diff(NewStringSet(v...)) }
func (s StringSet) Exists(v string) bool         { return len(s.Inter(StringSet{v})) != 0 }
