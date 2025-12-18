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

	if len(os.Args) < 2 {
		log.Fatalln("error: no filename specified")
	}

	total := 0
	filenames := os.Args[1:]

	for _, filename := range os.Args[1:] {
		wordCount := CountWordsInFile(filename)
		total = total + wordCount
		fmt.Println(wordCount, filename)
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

}

func CountWords(file io.Reader) int {
	countWords := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		countWords++
	}

	//**********
	// Using bufio.NewReader()
	// *********
	// isInsideWord := false

	// _ = isInsideWord

	// reader := bufio.NewReader(file)

	// for {
	// 	r, _, err := reader.ReadRune()

	// 	if err != nil {
	// 		break
	// 	}

	// 	if !unicode.IsSpace(r) && !isInsideWord {
	// 		countWords++
	// 	}

	// 	isInsideWord = !unicode.IsSpace(r)
	// }

	//**********
	// Doing it the long way
	// *********
	// const bufferSize = 2
	// buffer := make([]byte, bufferSize)
	// leftover := []byte{}

	// for {
	// 	size, err := file.Read(buffer)
	// 	if err != nil {
	// 		break
	// 	}

	// 	subBuffer := append(leftover, buffer[:size]...)
	// 	for len(subBuffer) > 0 {

	// 		r, rsize := utf8.DecodeRune(subBuffer)
	// 		if r == utf8.RuneError {
	// 			break
	// 		}
	// 		subBuffer = subBuffer[rsize:]

	// 		if !unicode.IsSpace(r) && !isInsideWord {
	// 			countWords++
	// 		}

	// 		isInsideWord = !unicode.IsSpace(r)
	// 	}

	// 	leftover = nil
	// 	leftover = append(leftover, subBuffer...)
	// }
	return countWords
}

func CountWordsInFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	return CountWords(file)
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
