package fuzzy

import (
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

func TestCanFuzzyMatchLazily(t *testing.T) {
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

		{
			fuzzy: "nomatch",
			input: words,
			want:  nil,
		},
	}

	t.Logf("call Matches() after each letter")
	for n, tt := range tests {
		matcher := LazyMatch(tt.input)
		for _, r := range []rune(tt.fuzzy) {
			matcher.Push(r)
			_, _ = matcher.matchesCost()
		}
		want := tt.want
		got, cost := matcher.matchesCost()

		t.Logf("test #%d, cost %d", n, cost)

		sort.Strings(want)
		sort.Strings(got)
		if !reflect.DeepEqual(want, got) {
			t.Logf("want=%#v", want)
			t.Logf(" got=%#v", got)
			t.Error("bad fuzzy match")
		}
	}

	t.Logf("call Matches() only at end")
	for n, tt := range tests {
		matcher := LazyMatch(tt.input)
		for _, r := range []rune(tt.fuzzy) {
			matcher.Push(r)
		}
		want := tt.want
		got, cost := matcher.matchesCost()

		t.Logf("test #%d, cost %d", n, cost)

		sort.Strings(want)
		sort.Strings(got)
		if !reflect.DeepEqual(want, got) {
			t.Logf("want=%#v", want)
			t.Logf(" got=%#v", got)
			t.Error("bad fuzzy match")
		}
	}
}

func TestCanFuzzyMatchEagerly(t *testing.T) {
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

		{
			fuzzy: "nomatch",
			input: words,
			want:  nil,
		},
	}

	t.Logf("call Matches() after each letter")
	for n, tt := range tests {

		matcher := EagerMatch(tt.input)
		for _, r := range []rune(tt.fuzzy) {
			matcher.Push(r)
			_, _ = matcher.matchesCost()
		}
		want := tt.want
		got, cost := matcher.matchesCost()

		t.Logf("test #%d, cost %d", n, cost)

		sort.Strings(want)
		sort.Strings(got)
		if !reflect.DeepEqual(want, got) {
			t.Logf("want=%#v", want)
			t.Logf(" got=%#v", got)
			t.Error("bad fuzzy match")
		}
	}

	t.Logf("call Matches() only at end")
	for n, tt := range tests {

		matcher := EagerMatch(tt.input)
		for _, r := range []rune(tt.fuzzy) {
			matcher.Push(r)
		}
		want := tt.want
		got, cost := matcher.matchesCost()

		t.Logf("test #%d, cost %d", n, cost)

		sort.Strings(want)
		sort.Strings(got)
		if !reflect.DeepEqual(want, got) {
			t.Logf("want=%#v", want)
			t.Logf(" got=%#v", got)
			t.Error("bad fuzzy match")
		}
	}

}
