package textstats

import (
	"reflect"
	"testing"
)

func TestCompute(t *testing.T) {
	cases := map[string]struct {
		in   string
		want Stats
	}{
		"Empty":             {in: "", want: Stats{}},
		"SingleLine":        {in: "hello world", want: Stats{Lines: 1, Words: 2, Chars: 11}},
		"TrailingNewline":   {in: "a b\nc\n", want: Stats{Lines: 2, Words: 3, Chars: 6}},
		"NoTrailingNewline": {in: "a b\nc", want: Stats{Lines: 2, Words: 3, Chars: 5}},
		"Unicode":           {in: "grüße", want: Stats{Lines: 1, Words: 1, Chars: 5}},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := Compute(tc.in); got != tc.want {
				t.Errorf("Compute(%q) = %+v, want %+v", tc.in, got, tc.want)
			}
		})
	}
}

func TestTopWords(t *testing.T) {
	text := "The cat and the hat. The CAT sat!"
	want := []WordFrequency{{"the", 3}, {"cat", 2}, {"and", 1}}

	if got := TopWords(text, 3); !reflect.DeepEqual(got, want) {
		t.Errorf("TopWords() = %+v, want %+v", got, want)
	}
	if got := TopWords(text, 0); got != nil {
		t.Errorf("TopWords(n=0) = %+v, want nil", got)
	}
	if got := TopWords("", 5); len(got) != 0 {
		t.Errorf("TopWords(empty) = %+v, want empty", got)
	}
}
