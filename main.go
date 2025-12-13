package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	data, err := os.ReadFile("./wors.txt")
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

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
