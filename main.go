package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	filename := "./words.txt"

	log.SetFlags(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	wordCount := CountWordsInFile(file)

	fmt.Println(wordCount)
}

func CountWordsInFile(file *os.File) int {
	countWords := 0
	isInsideWord := false

	const bufferSize = 7
	buffer := make([]byte, bufferSize)

	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		isInsideWord = !unicode.IsSpace(rune(buffer[0])) && isInsideWord

		bufferCount := CountWords(buffer[:size])
		if isInsideWord {
			bufferCount -= 1
		}
		countWords += bufferCount

		isInsideWord = !unicode.IsSpace(rune(buffer[size-1]))
	}
	return countWords
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
