package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	total := 0
	filenames := os.Args[1:]
	didError := false

	for _, filename := range os.Args[1:] {
		wordCount, err := CountWordsInFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "counter:", err)
			didError = true
			continue
		}
		total = total + wordCount
		fmt.Println(wordCount, filename)
	}

	if len(filenames) == 0 {
		wordCount := CountWords(os.Stdin)
		fmt.Println(wordCount)
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

	if didError {
		os.Exit(1)
	}
}

func CountWords(file io.Reader) int {
	countWords := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		countWords++
	}

	return countWords
}

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	count := CountWords(file)

	return count, nil
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
