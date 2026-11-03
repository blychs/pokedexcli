package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasur PIKACHU",
			expected: []string{"charmander", "bulbasur", "pikachu"},
		},
		{
			input: " hello     world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: " ",
			expected: []string{},
		},
		{
			input: "Hi",
			expected: []string{"hi"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match. len(c.expected): %v, len(actual): %v", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word: %s does not match expectedWord: %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
