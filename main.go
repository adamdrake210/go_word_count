package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")
	_ = data

	wordCount := 0
	const spaceChar = 32

	for i, x := range data {
		_ = i
		if x == spaceChar {
			wordCount++
		}
	}
	wordCount += 1

	fmt.Println(wordCount)
}
