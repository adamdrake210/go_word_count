package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	totals := Counts{}
	filenames := os.Args[1:]
	didError := false

	for _, filename := range os.Args[1:] {
		counts, err := CountFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "counter:", err)
			didError = true
			continue
		}
		totals = Counts{
			Bytes: totals.Bytes + counts.Bytes,
			Words: totals.Words + counts.Words,
			Lines: totals.Lines + counts.Lines,
		}
		counts.Print(os.Stdout, filename)
	}

	if len(filenames) == 0 {
		GetCounts(os.Stdin).Print(os.Stdout, "")
	}

	if len(filenames) > 1 {
		totals.Print(os.Stdout, "total")
	}

	if didError {
		os.Exit(1)
	}
}

// func CountWords(data []byte) int {
// 	wordCount := 0

// 	wasSpace := true
// 	for _, x := range data {
// 		isSpace := unicode.IsSpace(rune(x))
// 		if wasSpace && !isSpace {
// 			wordCount++
// 		}
// 		wasSpace = isSpace
// 	}

// 	return wordCount
// }
