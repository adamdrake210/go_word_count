package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	wants := 5

	result := CountWords([]byte(input))

	if result != wants {
		t.Logf("expected: %d got: %d", wants, result)
		t.Fail()
	}

	input = ""
	wants = 0

	result = CountWords([]byte(input))

	if result != wants {
		t.Logf("expected: %d got: %d", wants, result)
		t.Fail()
	}

	input = " "
	wants = 0

	result = CountWords([]byte(input))

	if result != wants {
		t.Logf("expected: %d got: %d", wants, result)
		t.Fail()
	}
}
