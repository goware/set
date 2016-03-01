package set_test

import (
	"reflect"
	"testing"

	"github.com/goware/set"
)

func TestInt64Set(t *testing.T) {
	s := set.NewInt64Set(3, 3, 1, 2)

	s = s.Add(5, 5, 3, 2, 1, 4, 4, 5)
	if len(s) != 5 {
		t.Error("expecting 5 values in the set")
	}
	if !reflect.DeepEqual([]int64(s), []int64{1, 2, 3, 4, 5}) {
		t.Error("invalid set")
	}

	s = s.Remove(5, 4)
	if len(s) != 3 {
		t.Error("expecting 3 values in the set")
	}
	if !reflect.DeepEqual([]int64(s), []int64{1, 2, 3}) {
		t.Error("invalid set")
	}

	s = s.Add(6, 1, 9, 13, 9, 8, 3)
	if len(s) != 7 {
		t.Error("expecting 7 values in the set, got:", len(s))
	}
	if !reflect.DeepEqual([]int64(s), []int64{1, 2, 3, 6, 8, 9, 13}) {
		t.Error("invalid set")
	}

	if !s.Exists(6) {
		t.Error("expecting value to exist in the set")
	}
	if s.Exists(100) {
		t.Error("expecting value to not exist in the set")
	}
	if !reflect.DeepEqual([]int64(s), []int64{1, 2, 3, 6, 8, 9, 13}) {
		t.Error("invalid set")
	}
}

func TestExistsInDescendingInt64Set(t *testing.T) {
	s := set.Int64Set{}
	s = s.Add(100)
	s = s.Add(10)
	s = s.Add(1)

	for _, v := range []int64{100, 10, 1} {
		if !s.Exists(v) {
			t.Errorf("expecting value %v to exist in the set", v)
		}
	}
	if !reflect.DeepEqual([]int64(s), []int64{1, 10, 100}) {
		t.Error("invalid set")
	}

}

func TestExistsInStringSet(t *testing.T) {
	words := []string{
		"www", "wwww", "mail", "api", "status", "cache", "cdn", "direct",
		"blog", "wiki", "faq", "help", "support", "knowledgebase", "mobile",
		"demo", "beta", "test", "testing", "prod", "production", "stg", "staging", "dev", "devel", "development",
		"root", "admin", "administrator", "superadmin", "sysadmin", "webmaster",
		"user", "account", "bot", "anonymous", "guest", "member", "collaborator", "moderator", "owner",
		"news", "job", "jobs", "career", "pricing", "about", "contact", "terms", "privacy", "tour",
		"search", "discover", "popular", "trending", "enterprise",
		"join", "invite", "register", "follow", "login", "logout", "auth", "oauth", "ping",
		"download", "upload", "app", "apps", "marketplace", "addons",
		"mention", "mentions", "tag", "tags", "email", "example", "old", "new",
	}

	s := set.NewStringSet(words...)

	for _, v := range words {
		if !s.Exists(v) {
			t.Errorf("expecting value %v to exist in the set", v)
		}
	}
}

func TestSubtractionInt64Set(t *testing.T) {
	first := set.Int64Set{}
	for _, i := range []int64{1, 2, 7, 9, 52, 53, 55, 63, 180, 212, 291, 307, 329, 365, 394, 395, 421, 448, 457, 484, 523, 533, 554, 561, 706, 896, 901, 906, 931, 959, 1081, 1086, 1088, 1090} {
		first = first.Add(i)
	}

	second := set.Int64Set{}
	for _, i := range []int64{1, 2, 7, 52, 55, 63, 180, 212, 291, 394, 395, 421, 448, 457, 533, 554, 561, 706, 901, 906, 1081, 1086, 1088, 1090} {
		second = second.Add(i)
	}

	firstMinusSecond := first.Remove(second...)

	for _, i := range second {
		if firstMinusSecond.Exists(i) {
			t.Errorf("Error! %v shouldn't be in firstMinusSecond", i)
		}
	}
}

func TestSubtractionInt64SetWithoutAdds(t *testing.T) {
	first := set.Int64Set{1, 2, 7, 9, 52, 53, 55, 63, 180, 212, 291, 307, 329, 365, 394, 395, 421, 448, 457, 484, 523, 533, 554, 561, 706, 896, 901, 906, 931, 959, 1081, 1086, 1088, 1090}

	second := set.Int64Set{1, 2, 7, 52, 55, 63, 180, 212, 291, 394, 395, 421, 448, 457, 533, 554, 561, 706, 901, 906, 1081, 1086, 1088, 1090}

	firstMinusSecond := first.Remove(second...)

	for _, i := range second {
		if firstMinusSecond.Exists(i) {
			t.Errorf("Error! %v shouldn't be in firstMinusSecond", i)
		}
	}
}

func TestStringSet(t *testing.T) {
	s := set.NewStringSet("a", "b", "a", "c")

	s = s.Add("d", "a", "b", "c", "e")
	if len(s) != 5 {
		t.Error("expecting 5 values in the set")
	}
	if !reflect.DeepEqual([]string(s), []string{"a", "b", "c", "d", "e"}) {
		t.Error("invalid set")
	}

	s = s.Remove("d", "a")
	if len(s) != 3 {
		t.Error("expecting 3 values in the set")
	}
	if !reflect.DeepEqual([]string(s), []string{"b", "c", "e"}) {
		t.Error("invalid set")
	}

	s = s.Add("z", "f", "e", "b", "c")
	if len(s) != 5 {
		t.Error("expecting 5 values in the set, got:", len(s))
	}
	if !reflect.DeepEqual([]string(s), []string{"b", "c", "e", "f", "z"}) {
		t.Error("invalid set")
	}

	if !s.Exists("b") {
		t.Error("expecting value to exist in the set")
	}
	if s.Exists("asdf") {
		t.Error("expecting value to not exist in the set")
	}

	if !reflect.DeepEqual([]string(s), []string{"b", "c", "e", "f", "z"}) {
		t.Error("invalid set")
	}

}
