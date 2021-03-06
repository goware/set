package set

import (
	"sort"

	"github.com/xtgo/set"
)

type BoolOp func(sort.Interface, int) bool

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
	sort.Sort(t)
	data := append(s, t...)
	n := op(data, len(s))
	return data[:n]
}

func (s Int64Set) DoBool(op BoolOp, t Int64Set) bool {
	data := append(s, t...)
	return op(data, len(s))
}

func (s Int64Set) Union(t Int64Set) Int64Set  { return s.Do(set.Union, t) }
func (s Int64Set) Inter(t Int64Set) Int64Set  { return s.Do(set.Inter, t) }
func (s Int64Set) Diff(t Int64Set) Int64Set   { return s.Do(set.Diff, t) }
func (s Int64Set) Add(v ...int64) Int64Set    { return s.Union(Int64Set(v)) }
func (s Int64Set) Remove(v ...int64) Int64Set { return s.Diff(NewInt64Set(v...)) }
func (s Int64Set) Exists(v int64) bool        { return s.DoBool(set.IsInter, Int64Set{v}) }

// String set
type StringSet []string

func NewStringSet(v ...string) StringSet {
	s := StringSet{}
	if len(v) > 0 {
		s = s.Add(sort.StringSlice(v)...)
	}
	return s
}

func (s StringSet) Do(op set.Op, t StringSet) StringSet {
	data := sort.StringSlice(append(s, t...))
	n := op(data, len(s))
	return StringSet(data[:n])
}

func (s StringSet) DoBool(op BoolOp, t StringSet) bool {
	data := sort.StringSlice(append(s, t...))
	return op(data, len(s))
}

func (s StringSet) Union(t StringSet) StringSet  { return s.Do(set.Union, t) }
func (s StringSet) Inter(t StringSet) StringSet  { return s.Do(set.Inter, t) }
func (s StringSet) Diff(t StringSet) StringSet   { return s.Do(set.Diff, t) }
func (s StringSet) Add(v ...string) StringSet    { return s.Union(StringSet(v)) }
func (s StringSet) Remove(v ...string) StringSet { return s.Diff(NewStringSet(v...)) }
func (s StringSet) Exists(v string) bool         { return s.DoBool(set.IsInter, StringSet{v}) }
