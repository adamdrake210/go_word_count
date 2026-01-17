package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func (c Counts) Print(w io.Writer, opts DisplayOptions, suffixes ...string) {
	xs := []string{}

	if opts.ShouldShowLines() {
		xs = append(xs, strconv.Itoa(c.Lines))
	}

	if opts.ShouldShowWords() {
		xs = append(xs, strconv.Itoa(c.Words))
	}

	if opts.ShouldShowBytes() {
		xs = append(xs, strconv.Itoa(c.Bytes))
	}

	xs = append(xs, suffixes...)

	line := strings.Join(xs, " ")

	//nolint
	fmt.Fprintln(w, line)
}

func GetCounts(file io.ReadSeeker) Counts {
	const offsetStart = 0

	bytes := CountBytes(file)
	_, _ = file.Seek(offsetStart, io.SeekStart)
	lines := CountLines(file)
	_, _ = file.Seek(offsetStart, io.SeekStart)
	words := CountWords(file)

	return Counts{Bytes: bytes, Lines: lines, Words: words}
}

func CountFile(filename string) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}

	//nolint:errcheck // intentionally ignoring error for defer
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
	byteCount, _ := io.Copy(io.Discard, r) //nolint:errcheck // intentionally ignoring error for discard
	return int(byteCount)
}
