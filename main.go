package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := "./words.txt"

	log.SetFlags(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	wordCount := CountWords(file)

	fmt.Println(wordCount)
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
