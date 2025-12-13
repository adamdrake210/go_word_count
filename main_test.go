package main_test

import (
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
			result := counter.CountWords([]byte(tc.input))
			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}
