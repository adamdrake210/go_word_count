package main_test

import (
	"bytes"
	"strings"
	"testing"

	counter "github.com/adamdrake210/word_counter"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "5 words with space at beginning",
			input: " one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "space empty",
			input: " ",
			wants: 0,
		},
		{
			name:  "new line",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "double space",
			input: "one two three.  Four five six",
			wants: 6,
		},
		{
			name:  "Suffix",
			input: "one two three.   ",
			wants: 3,
		},
		{
			name:  "Tab character",
			input: "Hello\tWord\n",
			wants: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := counter.GetCounts(strings.NewReader(tc.input)).Words

			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "simple five words, 1 new line",
			input: "one two three four five\n",
			wants: 1,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "no new lines",
			input: "one two three four five",
			wants: 0,
		},
		{
			name:  "no new lines at end",
			input: "one two three four five\nsix",
			wants: 1,
		},
		{
			name:  "multi newline string",
			input: "\n\n\n\n",
			wants: 4,
		},
		{
			name:  "multi word line string",
			input: "one\ntwo\nthree\nfour\nfive\n",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			res := counter.GetCounts(r).Lines

			if res != tc.wants {
				t.Logf("expected: %d, got %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "simple five words",
			input: "one two three four five",
			wants: 23,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "all spaces",
			input: "       ",
			wants: 7,
		},
		{
			name:  "newlines and words",
			input: "one\ntwo\nthree\nfour\nfive\t\n",
			wants: 25,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			numBytes := counter.GetCounts(r).Bytes

			if numBytes != tc.wants {
				t.Logf("expected: %d, got %d", tc.wants, numBytes)
				t.Fail()
			}
		})
	}
}

func TestGetCounts(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants counter.Counts
	}{
		{
			name:  "simple five words",
			input: "one two three four five\n",
			wants: counter.Counts{
				Lines: 1,
				Words: 5,
				Bytes: 24,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			res := counter.GetCounts(r)

			if res != tc.wants {
				t.Logf("expected: %d, got %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestPrintCounts(t *testing.T) {
	type inputs struct {
		counts   counter.Counts
		filename string
	}

	testCases := []struct {
		name  string
		input inputs
		wants string
	}{
		{
			name: "simple five words",
			input: inputs{
				counts: counter.Counts{

					Lines: 1,
					Words: 5,
					Bytes: 24,
				},
				filename: "words.txt",
			},
			wants: "1 5 24 words.txt\n",
		},
		{
			name: "no filename",
			input: inputs{
				counts: counter.Counts{
					Lines: 1,
					Words: 4,
					Bytes: 20,
				},
				filename: "",
			},
			wants: "1 4 20\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			tc.input.counts.Print(buffer, tc.input.filename)

			if buffer.String() != tc.wants {
				t.Logf("expected: %s, got %s", tc.wants, buffer.String())
				t.Fail()
			}
		})
	}
}
