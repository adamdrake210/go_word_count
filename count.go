package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountWords(file io.Reader) int {
	countWords := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		countWords++
	}

	return countWords
}

type Counts struct {
	Bytes int
	Words int
	Lines int
}

func (c Counts) Add(other Counts) Counts {
	c.Bytes += other.Bytes
	c.Lines += other.Lines
	c.Words += other.Words
	return c
}

func (c Counts) Print(w io.Writer, filenames ...string) {
	fmt.Fprintf(w, "%d %d %d", c.Lines, c.Words, c.Bytes)

	for _, filename := range filenames {
		fmt.Fprintf(w, " %s", filename)
	}

	fmt.Fprintf(w, "\n")
}

func GetCounts(file io.ReadSeeker) Counts {
	const offsetStart = 0

	bytes := CountBytes(file)
	file.Seek(offsetStart, io.SeekStart)
	lines := CountLines(file)
	file.Seek(offsetStart, io.SeekStart)
	words := CountWords(file)

	return Counts{Bytes: bytes, Lines: lines, Words: words}
}

func CountFile(filename string) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}

	defer file.Close()

	return GetCounts(file), nil
}

func CountLines(r io.Reader) int {
	countLines := 0

	reader := bufio.NewReader(r)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if r == '\n' {
			countLines++
		}
	}

	return countLines
}

func CountBytes(r io.Reader) int {
	byteCount, _ := io.Copy(io.Discard, r)
	return int(byteCount)
}
