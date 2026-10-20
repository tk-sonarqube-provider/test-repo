// Package textstats computes simple statistics over a block of text.
package textstats

import (
	"sort"
	"strings"
	"unicode"
)

// Stats summarises a block of text.
type Stats struct {
	Lines int
	Words int
	Chars int
}

// Compute returns line, word and character counts for text.
func Compute(text string) Stats {
	if text == "" {
		return Stats{}
	}
	return Stats{
		Lines: strings.Count(text, "\n") + boolToInt(!strings.HasSuffix(text, "\n")),
		Words: len(strings.Fields(text)),
		Chars: len([]rune(text)),
	}
}

// WordFrequency is a word and how often it occurs.
type WordFrequency struct {
	Word  string
	Count int
}

// TopWords returns the n most frequent words in text, lower-cased and stripped
// of surrounding punctuation. Ties are broken alphabetically.
func TopWords(text string, n int) []WordFrequency {
	if n <= 0 {
		return nil
	}

	counts := map[string]int{}
	for _, field := range strings.Fields(text) {
		word := strings.ToLower(strings.TrimFunc(field, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		}))
		if word != "" {
			counts[word]++
		}
	}

	freqs := make([]WordFrequency, 0, len(counts))
	for w, c := range counts {
		freqs = append(freqs, WordFrequency{Word: w, Count: c})
	}
	sort.Slice(freqs, func(i, j int) bool {
		if freqs[i].Count != freqs[j].Count {
			return freqs[i].Count > freqs[j].Count
		}
		return freqs[i].Word < freqs[j].Word
	})

	if len(freqs) > n {
		freqs = freqs[:n]
	}
	return freqs
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
