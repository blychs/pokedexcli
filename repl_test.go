package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello",
			expected: []string{"hello"},
		},
		{
			input:    " HELLO   world!",
			expected: []string{"hello", "world!"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "Hello world, you are great",
			expected: []string{"hello", "world,", "you", "are", "great"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match, actual %d, got %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words don't match, expected %s, got %s", expectedWord, word)
			}
		}
	}
}

func TestCommandRegistry(t *testing.T) {
}
