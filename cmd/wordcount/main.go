// Command wordcount prints line, word and character counts for stdin, plus the
// most frequent words.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/tk-sonarqube-provider/test-repo/internal/textstats"
)

func main() {
	top := flag.Int("top", 5, "number of most frequent words to print")
	flag.Parse()

	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read stdin:", err)
		os.Exit(1)
	}

	text := string(in)
	s := textstats.Compute(text)
	fmt.Printf("lines=%d words=%d chars=%d\n", s.Lines, s.Words, s.Chars)
	for _, wf := range textstats.TopWords(text, *top) {
		fmt.Printf("%6d %s\n", wf.Count, wf.Word)
	}
}
