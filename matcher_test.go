package fuzzy_test

import (
	"github.com/aybabtme/fuzzy"
	"reflect"
	"sort"
	"testing"
)

var words = []string{
	"reprecipitation",
	"grallic",
	"fir",
	"emigrate",
	"cataphrenia",
	"unpostponed",
	"prerogativity",
	"chiefly",
	"hup",
	"unzealously",
	"goldilocks",
	"especial",
	"exoticness",
	"polymorphean",
	"chalcosine",
	"tutworkman",
	"labrosaurid",
	"compactness",
	"superannuate",
	"uranist",
}

func TestCanFuzzyMatch(t *testing.T) {
	tests := []struct {
		fuzzy string
		input []string
		want  []string
	}{
		{
			fuzzy: "grallic",
			input: words,
			want:  []string{"grallic"},
		},
		{
			fuzzy: "z",
			input: words,
			want:  []string{"unzealously"},
		},
		{
			fuzzy: "g",
			input: words,
			want:  []string{"emigrate", "goldilocks", "grallic", "prerogativity"},
		},

		{
			fuzzy: "ga",
			input: words,
			want:  []string{"emigrate", "grallic", "prerogativity"},
		},
		{
			fuzzy: "gai",
			input: words,
			want:  []string{"grallic", "prerogativity"},
		},
		{
			fuzzy: "gaiy",
			input: words,
			want:  []string{"prerogativity"},
		},
		{
			fuzzy: "em",
			input: words,
			want:  []string{"emigrate"},
		},
		{
			fuzzy: "te",
			input: words,
			want:  []string{"cataphrenia", "compactness", "emigrate", "exoticness", "superannuate", "unpostponed"},
		},

		{
			fuzzy: "tes",
			input: words,
			want:  []string{"compactness", "exoticness"},
		},
	}

	for n, tt := range tests {
		t.Logf("test #%d", n)
		matcher := fuzzy.Match(tt.input)
		for _, r := range []rune(tt.fuzzy) {
			matcher.Push(r)
		}
		want := tt.want
		got := matcher.Matches()
		sort.Strings(want)
		sort.Strings(got)
		if !reflect.DeepEqual(want, got) {
			t.Logf("want=%#v", want)
			t.Logf(" got=%#v", got)
			t.Error("bad fuzzy match")
		}
	}
}
