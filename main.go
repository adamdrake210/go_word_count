package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := CountWords(data)

	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	wordCount := 0

	wasSpace := true
	for _, x := range data {
		isSpace := unicode.IsSpace(rune(x))
		if wasSpace && !isSpace {
			wordCount++
		}
		wasSpace = isSpace
	}

	return wordCount
}
